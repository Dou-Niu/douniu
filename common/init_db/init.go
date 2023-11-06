package init_db

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"

	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// gorm初始化
var DB *gorm.DB

func InitGorm(MysqlDataSourece string) *gorm.DB {
	// 将日志写进kafka
	// logx.SetWriter(*LogxKafka())
	var err error
	DB, err = gorm.Open(mysql.Open(MysqlDataSourece),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
			Logger: logger.Default.LogMode(logger.Silent),
		})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	db, _ := DB.DB()
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(500)
	db.SetConnMaxLifetime(time.Minute)
	return DB
}

// redis初始化
func InitRedis(addr, password string) *redis.Client {
	// 创建一个 Redis 客户端连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis 服务器地址
		Password: password, // 如果有密码的话
		DB:       0,        // 默认数据库
	})

	// 使用 Ping 操作检查连接是否正常
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("连接redis失败, error=" + err.Error())
	}
	fmt.Println("redis连接成功")

	return rdb
}

// minio初始化
func InitMinio(Endpoint, AccessKey, SecretKey string) *minio.Client {
	// 连接到 MinIO 集群
	endpoint := Endpoint
	accessKey := AccessKey
	secretKey := SecretKey
	useSSL := false

	// 创建一个MinIO客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logx.Error("连接minio失败, error=", err)
		panic("连接minio失败, error=" + err.Error())
	}
	fmt.Println("连接minio集群成功")
	return minioClient
}
func GetESClient(url string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		panic(err)
	}
	fmt.Println("ES initialized...")

	return client

}
