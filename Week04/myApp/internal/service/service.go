package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	s *http.Server
}

func New() *Server {
	route := gin.New()
	route.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world")
	})
	s := &http.Server{
		Addr:    ":8888",
		Handler: route,
	}
	return &Server{s: s}

}

func (s *Server) Start() {
	if err := s.s.ListenAndServe(); err != nil {
		log.Fatalf("err: %v", err)
	}

}
