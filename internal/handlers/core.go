package handlers

import (
	"github.com/gin-gonic/gin"
)

type srv struct {
	engine *gin.Engine
}

func New() *srv {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Next()
	})
	return &srv{engine: r}
}

func (s *srv) Router() *gin.RouterGroup {
	return s.engine.Group("/api/")
}

func (s *srv) Run() error {
	return s.engine.Run()
}
