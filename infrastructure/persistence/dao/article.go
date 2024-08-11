package dao

import (
	"context"
	"gorm.io/gorm"
)

type Article struct {
	Id          int64  `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255)"`
	ContentMD   string
	ContentHTML string
	AuthorId    int64 `gorm:"index"`
	CreatedAt   int64
	UpdatedAt   int64
}

type ArticleDao interface {
	Insert(ctx context.Context, a *Article) (int64, error)
	Update(ctx context.Context, a *Article) error
	FindById(ctx context.Context, id int64) (Article, error)
	FindByAuthorId(ctx context.Context, authorId int64) ([]*Article, error)
}

type GORMArticleDao struct {
	db *gorm.DB
}

func (dao *GORMArticleDao) Insert(ctx context.Context, a *Article) (int64, error) {
	err := dao.db.Create(a).Error
	if err != nil {
		return 0, err
	}
	return a.Id, nil
}

func (dao *GORMArticleDao) Update(ctx context.Context, a *Article) error {
	return dao.db.Save(a).Error
}

func (dao *GORMArticleDao) FindById(ctx context.Context, id int64) (Article, error) {
	var a Article
	err := dao.db.Where("id = ?", id).First(&a).Error
	if err != nil {
		return a, err
	}
	return a, nil
}

func (dao *GORMArticleDao) FindByAuthorId(ctx context.Context, authorId int64) ([]*Article, error) {
	var as []*Article
	err := dao.db.Where("author_id = ?", authorId).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}
