package user

import "time"

type User struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name string) *User {
	return &User{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
