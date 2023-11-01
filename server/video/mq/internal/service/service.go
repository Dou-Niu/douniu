package service

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/init_db"
	"douniu/server/common/utils"
	"douniu/server/video/model"
	"douniu/server/video/mq/internal/config"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"

	"github.com/zeromicro/go-zero/core/jsonx"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	chanCount   = 10   // 通道数量
	bufferCount = 1024 // 每个通道的缓冲区大小
)

type Service struct {
	c             config.Config       // 配置信息
	waiter        sync.WaitGroup      // 用于等待所有消费者 goroutine 完成的等待组
	msgsChan      []chan *model.Video // 消息通道切片，每个元素是一个通道，用于存放消息
	SensitiveTrie *utils.SensitiveTrie
	SqlConn       sqlx.SqlConn
	RedisClient   *redis.Client
	VideoModel    model.VideoModel
	ctx           context.Context
	ESClient      *elastic.Client
}

// NewService 创建一个新的 Service 实例
func NewService(c config.Config) *Service {
	mysqlConn := sqlx.NewMysql(c.MysqlConf.DataSource)

	// 创建 Service 实例
	s := &Service{
		c:             c,
		ctx:           context.Background(),
		msgsChan:      make([]chan *model.Video, chanCount),
		SqlConn:       mysqlConn,
		VideoModel:    model.NewVideoModel(mysqlConn, c.CacheRedis),
		SensitiveTrie: utils.NewSensitiveTrie(),
		RedisClient:   init_db.InitRedis(c.RedisConf.Host, c.RedisConf.Password),
		ESClient:      init_db.GetESClient(c.ESConf.Host),
	}

	// 创建 chanCount 个消费者 goroutine
	for i := 0; i < chanCount; i++ {
		ch := make(chan *model.Video, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}

	return s
}

// consume 是消费者 goroutine 的函数，负责处理从通道中接收的消息
func (s *Service) consume(ch chan *model.Video) {
	defer s.waiter.Done()

	for {
		message, ok := <-ch
		if !ok {
			log.Fatal("接受消息失败")
		}
		m := *message
		fmt.Printf("消费消息: %+v\n", m)

		// 创建 video 对象，用于写入数据库
		v := model.Video{
			Id:         m.Id,
			UserId:     m.UserId,
			Title:      m.Title,
			PlayUrl:    m.PlayUrl,
			CoverUrl:   m.CoverUrl,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			DeleteTime: time.Now(),
			Partition:  m.Partition,
		}
		// 敏感词过滤
		s.SensitiveTrie.AddWords([]string{"傻逼", "死", "你妈", "滚"})
		v.Title = s.SensitiveTrie.Filter(v.Title)
		fmt.Println(v)
		// 并发写入redis,mysql
		err := mr.Finish(func() error {
			// 写入mysql
			_, err := s.VideoModel.Insert(s.ctx, &v)
			return err
		}, func() error {
			// 全部视频时序排序
			err := s.RedisClient.ZAdd(s.ctx, consts.VideoTimeScore, redis.Z{
				Score:  float64(v.CreateTime.Unix()),
				Member: m.Id,
			}).Err()

			return err
		}, func() error {
			// 全部视频热度排序
			err := s.RedisClient.ZAdd(s.ctx, consts.VideoHotScore, redis.Z{
				Score:  float64(v.CreateTime.Unix()) * 0.5,
				Member: m.Id,
			}).Err()
			return err
		}, func() error {
			// 用户所有视频时序排序
			err := s.RedisClient.ZAdd(s.ctx, consts.VideoEveryUserTimeScore+fmt.Sprint(m.UserId), redis.Z{
				Score:  float64(v.CreateTime.Unix()),
				Member: m.Id,
			}).Err()
			return err
		}, func() error {
			// 用户所有视频热度排序
			err := s.RedisClient.ZAdd(s.ctx, consts.VideoEveryUserHotScore+fmt.Sprint(m.UserId), redis.Z{
				Score:  float64(v.CreateTime.Unix()) * 0.5,
				Member: m.Id,
			}).Err()
			return err
		}, func() error {
			// 分区所有视频时序排序
			err := s.RedisClient.ZAdd(s.ctx, consts.VideoPartitionTimeScore+fmt.Sprint(m.Partition), redis.Z{
				Score:  float64(v.CreateTime.Unix()),
				Member: m.Id,
			}).Err()
			return err
		}, func() error {
			// 分区所有视频热度排序
			err := s.RedisClient.ZAdd(s.ctx, consts.VideoPartitionHotScore+fmt.Sprint(m.Partition), redis.Z{
				Score:  float64(v.CreateTime.Unix()) * 0.5,
				Member: m.Id,
			}).Err()
			return err
		})
		if err != nil {
			logx.Error("video mq并发写入mysql,redis错误，err:", err)
			return
		}
		vJson, err := jsonx.MarshalToString(v)
		if err != nil {
			logx.Error("解析json对象，err:", err)

			return
		}
		_, err = s.ESClient.Index().
			Index(consts.EsVideoIndex).
			BodyJson(vJson).
			Type(consts.EsVideoIndex).
			Do(s.ctx)
		if err != nil {
			logx.Error("video同步到es错误，err:", err)
			return
		}

		logx.Infof("video= %v 写入成功", m.Id)
	}
}

// Consume 是消费者的方法，用于处理消息
func (s *Service) Consume(_, value string) error {
	logx.Infof("消费消息: %s\n", value)

	// 将 JSON 数据解析为 []*model.NewUserFile 对象
	var data []*model.Video
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	// 将解析后的消息根据 UserId 分发到不同的通道
	for _, d := range data {
		s.msgsChan[d.Id%chanCount] <- d
	}

	return nil
}
