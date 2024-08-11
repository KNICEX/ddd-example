package cache

import (
	"context"
	"github.com/KNICEX/DDD-example/domain/article"
)

type ArticleCache interface {
	Get(ctx context.Context, id int64) (article.Article, error)
	Set(ctx context.Context, id int64, art *article.Article) error
}
