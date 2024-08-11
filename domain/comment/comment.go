package comment

import "time"

type Comment struct {
	Id        int64
	Content   string
	CreatedAt time.Time

	Article Article
	Author  Author
}

type Article struct {
	Id int64
}

type Author struct {
	Id   int64
	Name string
}
