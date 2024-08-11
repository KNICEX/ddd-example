package article

import (
	"context"
)

type Repository interface {
	// Create 创建后将Id设置回来
	Create(ctx context.Context, article *Article) error
	Update(ctx context.Context, article *Article) error
	FindByID(ctx context.Context, id int64) (*Article, error)
	FindByAuthor(ctx context.Context, authorID int64) ([]*Article, error)
}
