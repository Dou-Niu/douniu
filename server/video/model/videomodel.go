package model

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/errorx"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"time"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindByTimeOrHot(ctx context.Context, rdb *redis.Client, key string, max int64) ([]int64, error)
		FindFollowFeed(ctx context.Context, followIds []int64, leastTime time.Time) ([]int64, error)
		GetZsetCount(ctx context.Context, rdb *redis.Client, key string) (int64, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

func (m *customVideoModel) GetZsetCount(ctx context.Context, rdb *redis.Client, key string) (int64, error) {
	// 使用ZCARD命令获取有序集合的元素个数
	count, err := rdb.ZCard(ctx, key).Result()
	if err != nil {
		logx.Error("获取有序集合元素个数出错:", err)
		return 0, err
	}
	return count, nil
}

func (m *customVideoModel) FindFollowFeed(ctx context.Context, followIds []int64, leastTime time.Time) ([]int64, error) {
	placeholders := make([]interface{}, len(followIds))
	for i, id := range followIds {
		placeholders[i] = id
	}

	query := fmt.Sprintf("select id from %s where user_id IN (?) and create_time < ? order by create_time DESC limit %v", m.table, consts.DefaultSizeOfPage)
	placeholders = append(placeholders, leastTime)

	var resp []int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, placeholders...)
	if len(resp) == 0 {
		logx.Error(err)
		err = errors.New("查不到数据了")
	}
	return resp, err
}

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c, opts...),
	}
}
func (m *customVideoModel) FindByTimeOrHot(ctx context.Context, rdb *redis.Client, key string, max int64) ([]int64, error) {
	// 使用ZREVRANGEBYSCORE命令查询score小于max的最大10个成员
	vIds, err := rdb.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min: "0",
		// -1防止重复
		Max:    fmt.Sprint(max - 1),
		Offset: 0,
		Count:  consts.DefaultSizeOfPage,
	}).Result()
	if err != nil {
		logc.Error(ctx, err)

		return nil, errors.Wrapf(errorx.NewDefaultError("redis ZREVRANGEBYSCORE错误"+err.Error()), "redis ZREVRANGEBYSCORE错误%v", err)
	}
	if len(vIds) == 0 {
		// 没有找到匹配的成员，处理逻辑
		return nil, errors.Wrapf(errorx.NewDefaultError("没有下一页的数据了"), "")

	}

	res := make([]int64, 0)
	for _, v := range vIds {
		v64, _ := strconv.ParseInt(v.Member.(string), 10, 64)
		res = append(res, v64)
	}
	return res, nil
}
