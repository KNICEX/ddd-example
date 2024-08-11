package article

import "time"

type Article struct {
	Id          int64
	Title       string
	ContentMD   string
	ContentHTML string
	Author      Author
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Author struct {
	Id   int64
	Name string
}

func NewArticle(title, contentMD string, author Author) *Article {
	// 创建文章
	a := &Article{
		Title:     title,
		ContentMD: contentMD,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	a.md2html()
	return a
}

func (a *Article) md2html() {
	// TODO 转换内容, escape...
	a.ContentHTML = a.ContentMD
}

func (a *Article) Update(title, content string) {
	// TODO 校验内容
	a.Title = title
	a.ContentMD = content
	a.md2html()
}
