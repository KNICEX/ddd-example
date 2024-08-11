package application

import (
	"context"
	"github.com/KNICEX/DDD-example/domain/article"
)

type ArticleService interface {
	Save(ctx context.Context, articleId int64, title string, contentMD string, authorId int64) (int64, error)
	Publish(ctx context.Context, articleId int64) error
}

type articleService struct {
	svc  article.Service
	repo article.Repository
}

func NewArticleService(repo article.Repository) ArticleService {
	return &articleService{
		repo: repo,
	}
}

func (a *articleService) Save(ctx context.Context, articleId int64, title string, contentMD string, authorId int64) (int64, error) {
	art := article.NewArticle(title, contentMD, article.Author{
		Id: authorId,
	})
	art.Id = articleId

	// 如果 articleId 为 0，创建一个新草稿
	if articleId == 0 {
		err := a.repo.Create(ctx, art)
		return art.Id, err
	}
	// 如果 articleId 不为 0，更新草稿
	return art.Id, a.repo.Update(ctx, art)
}

func (a *articleService) Publish(ctx context.Context, articleId int64) error {
	return a.svc.PublishArticle(&article.Article{
		Id: articleId,
	})
}
