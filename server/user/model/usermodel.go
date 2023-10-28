package model

import (
	"context"
	"douniu/server/common/consts"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		IsCodeVerify(ctx context.Context, rdb *redis.Client, phone string, code string) error
		ChangePassword(ctx context.Context, userId int64, newPassword string) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (m *customUserModel) ChangePassword(ctx context.Context, userId int64, newPassword string) error {
	one, err := m.FindOne(ctx, userId)
	if err != nil {
		return err
	}
	one.Password = newPassword
	err = m.Update(ctx, one)

	return err
}

func (m *customUserModel) IsCodeVerify(ctx context.Context, rdb *redis.Client, phone string, code string) error {

	TrueCode, err := rdb.Get(ctx, consts.PhoneCode+phone).Result()
	if err != nil || TrueCode != code {
		return errors.New("验证码错误")
	}
	return nil
}

func (m *customUserModel) IsRegister(ctx context.Context, rdb *redis.Client, phone string) (is bool, err error) {
	isExists, err := rdb.Exists(ctx, consts.UserIsRegister+phone).Result()
	if err != nil {
		return false, err
	}
	return isExists == 1, nil
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}
