package post

import "gorm.io/gorm"

type Postgres struct {
	DB *gorm.DB
}

func (p Postgres) SavePost(post Post) {
	p.DB.Save(post)
}
