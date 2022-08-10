package post

import "time"

type Post struct {
	ID          string
	Title       string
	Description string
	Created_at  time.Time
	Updated_at  time.Time
	Deleted_at  time.Time
}

type Store interface {
	Save(post Post) error
}
