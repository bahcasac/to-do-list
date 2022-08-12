package postgres

import (
	"context"
	"fmt"
	"github.com/bahcasac/stock/post"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

//func TestCRUD(t *testing.T) {
//	ctx := context.Background()
//	container, err := person.SetupMysqL(ctx)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer container.Terminate(ctx)
//	db, err := sql.Open("mysql", container.URI)
//	if err != nil {
//		t.Error(err)
//	}
//	defer db.Close()
//	err = person.InitMySQL(ctx, db)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer person.TruncateMySQL(ctx, db)
//
//	repo := mysql.NewMySQL(db)
//
//	t.Run("inserir person", func(t *testing.T) {
//		p := &person.Person{
//			Name:     "Ozzy",
//			LastName: "Osbourne",
//		}
//		id, err := repo.Create(p)
//		assert.Equal(t, person.ID(1), id)
//		assert.Nil(t, err)
//	})

func TestCRUD(t *testing.T) {
	container, err := post.SetupPostgres(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer container.Close(t)

	fmt.Println("passou aqui...")
	db, err := connection(container.DSN)
	if err != nil {
		t.Fatal(err)
	}
	defer container.Close(t)

	p := Postgres{
		DB: db,
	}
	post := post.Post{
		ID:          "124",
		Title:       "Teste 02",
		Description: "Este é um exemplo de post",
		Created_at:  time.Now(),
		Updated_at:  time.Time{},
		Deleted_at:  time.Time{},
	}
	err = p.SavePost(post)
	assert.NoError(t, err)
}

func connection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		println("aqui", dsn)
		return nil, err
	}
	return db, nil
}
