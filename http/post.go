package http

import (
	"fmt"
	"github.com/bahcasac/stock/http/presenter"
	"github.com/bahcasac/stock/post"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(s post.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		var post *post.Post

		if err := c.BindJSON(&post); err != nil {

			c.String(http.StatusBadRequest, fmt.Sprintf("error handling with body: %v", err))
			return
		}

		err := s.Create(*post)
		if err != nil {
			fmt.Sprintf("[Handler] CreateProducts error couldn't create products: %v", err)
			return
		}

		postResponse := presenter.CreatePostResponse{}

		c.JSON(http.StatusOK, postResponse)
	}
}

func MakePostHandler(r *gin.RouterGroup, s post.UseCase) {
	r.Handle("POST", "/post", CreatePost(s))

}
