package article

type Service interface {
	PublishArticle(article *Article) error
}

type articleService struct {
	repo Repository
}

func NewService() Service {
	return &articleService{}
}

func (a *articleService) PublishArticle(article *Article) error {
	// TODO 发布文章，将记录写入到公开库/表

	// TODO 其他操作， 比如发送通知, 更新用户文章数等
	return nil
}
