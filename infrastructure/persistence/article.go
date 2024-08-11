package persistence

import (
	"context"
	"github.com/KNICEX/DDD-example/domain/article"
	"github.com/KNICEX/DDD-example/infrastructure/persistence/cache"
	"github.com/KNICEX/DDD-example/infrastructure/persistence/dao"
	"time"
)

var _ article.Repository = (*CachedArticleRepo)(nil)

type CachedArticleRepo struct {
	dao   dao.ArticleDao
	cache cache.ArticleCache
}

func NewCachedArticleRepo(dao dao.ArticleDao) article.Repository {
	return &CachedArticleRepo{
		dao: dao,
	}
}

func (c *CachedArticleRepo) Create(ctx context.Context, article *article.Article) error {
	now := time.Now()
	article.CreatedAt = now
	article.UpdatedAt = now
	id, err := c.dao.Insert(ctx, c.domainToEntity(article))
	article.Id = id
	return err
}

func (c *CachedArticleRepo) Update(ctx context.Context, article *article.Article) error {
	article.UpdatedAt = time.Now()
	return c.dao.Update(ctx, c.domainToEntity(article))
}

func (c *CachedArticleRepo) FindByID(ctx context.Context, id int64) (*article.Article, error) {
	// TODO Check cache

	art, err := c.dao.FindById(ctx, id)

	return c.entityToDomain(&art), err
}

func (c *CachedArticleRepo) FindByAuthor(ctx context.Context, authorID int64) ([]*article.Article, error) {
	arts, err := c.dao.FindByAuthorId(ctx, authorID)

	var articles []*article.Article
	for _, art := range arts {
		articles = append(articles, c.entityToDomain(art))
	}

	return articles, err
}

func (c *CachedArticleRepo) domainToEntity(a *article.Article) *dao.Article {
	return &dao.Article{
		Id:          a.Id,
		Title:       a.Title,
		ContentMD:   a.ContentMD,
		ContentHTML: a.ContentHTML,
		AuthorId:    a.Author.Id,
		CreatedAt:   a.CreatedAt.UnixMilli(),
		UpdatedAt:   a.UpdatedAt.UnixMilli(),
	}
}

func (c *CachedArticleRepo) entityToDomain(a *dao.Article) *article.Article {
	return &article.Article{
		Id:          a.Id,
		Title:       a.Title,
		ContentMD:   a.ContentMD,
		ContentHTML: a.ContentHTML,
		Author: article.Author{
			Id: a.AuthorId,
		},
		CreatedAt: time.UnixMilli(a.CreatedAt),
		UpdatedAt: time.UnixMilli(a.UpdatedAt),
	}
}
