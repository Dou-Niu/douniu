package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
	FollowModel interface {
		followModel
		FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Follow, error)
		DeleteByUIdAndFId(ctx context.Context, userId int64, followId int64) error
		FindFriendIdList(ctx context.Context, userId int64) ([]int64, error)
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

func (m *customFollowModel) FindFriendIdList(ctx context.Context, userId int64) ([]int64, error) {
	query := `SELECT DISTINCT f1.follow_id FROM follow f1
         INNER JOIN follow f2 ON f1.follow_id = f2.user_id
		WHERE f1.user_id = ? AND f2.follow_id = ? order by update_time desc ;`
	idList := make([]int64, 0)
	err := m.conn.QueryRowsCtx(ctx, &idList, query, userId, userId)
	return idList, err
}

func (m *customFollowModel) DeleteByUIdAndFId(ctx context.Context, userId int64, followId int64) error {
	query := fmt.Sprintf("delete from %s where user_id = ? and follow_id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId, followId)
	return err
}

func (m *customFollowModel) FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Follow, error) {
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", followRows, m.table)
	var resp []*Follow
	err := m.conn.QueryRowsCtx(ctx, &resp, query, (pageNum-1)*pageSize, pageSize)

	return resp, err
}

// NewFollowModel returns a model for the database table.
func NewFollowModel(conn sqlx.SqlConn) FollowModel {
	return &customFollowModel{
		defaultFollowModel: newFollowModel(conn),
	}
}
