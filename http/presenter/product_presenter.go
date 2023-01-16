package presenter

import (
	"github.com/bahcasac/stock/post"
)

type PostPresenter struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreatePostResponse struct {
	Post         *post.Post `json:"post"`
	ErrorMessage string     `json:"errorMessage"`
}
