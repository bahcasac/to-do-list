package postgres

import (
	"github.com/bahcasac/stock/post"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (p Postgres) SavePost(post post.Post) error {
	result := p.DB.Save(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
