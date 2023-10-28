package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Comment, error)
		FindByVideoId(ctx context.Context, videoId int64) ([]*Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

func (m *customCommentModel) FindByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	query := fmt.Sprintf("select %s from %s where video_id = ?", commentRows, m.table)
	var resp []*Comment
	err := m.conn.QueryRowsCtx(ctx, &resp, query, videoId)
	return resp, err
}

func (m *customCommentModel) FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Comment, error) {
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", commentRows, m.table)
	var resp []*Comment
	err := m.conn.QueryRowsCtx(ctx, &resp, query, (pageNum-1)*pageSize, pageSize)

	return resp, err
}

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn),
	}
}
