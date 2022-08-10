package http

// request := gin.Default()
// request.GET("/helth", func(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"up": "Server is up and running baby!",
// 	})
// })
// request.Run()

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

// NewServer return a new fiber server
func NewServer() (*Server, error) {
	server := &Server{}
	r := gin.Default()
	server.httpServer = &http.Server{
		Addr:    ":8087",
		Handler: r,
	}
	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return server, nil
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
