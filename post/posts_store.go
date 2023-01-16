package post

import "gorm.io/gorm"

type Postgres struct {
	DB *gorm.DB
}

type Params struct {
	field, condition, value interface{}
}

func (p Postgres) SavePost(post Post) error {
	result := p.DB.Save(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p Postgres) FindAll() ([]Post, error) {
	var collection []Post
	result := p.DB.Find(&collection)
	if result.Error != nil {
		return []Post{}, result.Error
	}
	return collection, nil
}

func (p Postgres) Find() ([]Post, error) {
	var collection []Post
	result := p.DB.Find(&collection)
	if result.Error != nil {
		return []Post{}, result.Error
	}
	return collection, nil
}
