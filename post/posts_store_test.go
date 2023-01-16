package post

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

func TestCRUD(t *testing.T) {
	container := NewTestDatabase()
	db, err := connection(container.ConnectionString(t))
	if err != nil {
		log.Fatal(err)
	}

	p := Postgres{
		DB: db,
	}
	item := Post{
		ID:          "124",
		Title:       "Teste 02",
		Description: "Este é um exemplo de post",
		Created_at:  time.Now(),
		Updated_at:  time.Time{},
		Deleted_at:  time.Time{},
	}

	err = container.AutoMigrate(db, item)
	if err != nil {
		log.Fatal(err)
	}
	err = p.SavePost(item)
	container.Flush(db)
	container.Close(t)
	assert.NoError(t, err)
}

func connection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
