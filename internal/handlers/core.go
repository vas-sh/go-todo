package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

type userFetcher interface {
	GetUser(auth string) (*models.User, error)
}

type srv struct {
	engine  *gin.Engine
	userSrv userFetcher
}

func New(userSrv userFetcher) *srv {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Next()
	})
	return &srv{
		engine:  r,
		userSrv: userSrv,
	}
}

func (s *srv) AnonRouter() *gin.RouterGroup {
	return s.engine.Group("/api/")
}

func (s *srv) AuthRouter() *gin.RouterGroup {
	router := s.AnonRouter()
	router.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		user, err := s.userSrv.GetUser(c.GetHeader("Authorization"))
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set(models.UserContextKey, user)
		c.Next()
	})
	return router
}

func (s *srv) Run() error {
	return s.engine.Run()
}
