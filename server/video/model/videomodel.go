package model

import (
	"context"
	"douniu/server/common/consts"
	"douniu/server/common/errorx"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindByTimeOrHot(ctx context.Context, rdb *redis.Client, key string, max int64) ([]int64, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

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
