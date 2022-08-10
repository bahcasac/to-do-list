package main

import (
	"fmt"
	"github.com/bahcasac/stock/config"
	"github.com/bahcasac/stock/post"
	"time"
)

func main() {
	db, err := config.ConnectionDB()
	if err != nil {
		fmt.Println(err)
	}

	p := post.Postgres{
		DB: db,
	}
	post := post.Post{
		ID:          "123",
		Title:       "Teste 02",
		Description: "Este é um exemplo de post",
		Created_at:  time.Now(),
		Updated_at:  time.Time{},
		Deleted_at:  time.Time{},
	}
	p.SavePost(post)

}
