package main

import (
	"fmt"
	"github.com/bahcasac/stock/config"
	"github.com/bahcasac/stock/post"
	"github.com/bahcasac/stock/post/postgres"
	"log"
	"time"
)

func main() {
	db, err := config.ConnectionDB()
	if err != nil {
		fmt.Println(err)
	}

	p := postgres.Postgres{
		DB: db,
	}

	err = p.SavePost(post.Post{
		ID:          "123",
		Title:       "Teste 02",
		Description: "Este é um exemplo de post",
		Created_at:  time.Now(),
		Updated_at:  time.Time{},
		Deleted_at:  time.Time{},
	})
	if err != nil {
		log.Fatal(err)
	}

}
