package wshandlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/vas-sh/todo/internal/models"
)

type userFetcher interface {
	GetUser(auth string) (*models.User, error)
}

type srv struct {
	sync.RWMutex
	allConn     map[int64][]userConn
	userFetcher userFetcher
	pingTimeout time.Duration
	upgrader    websocket.Upgrader
}

func New(userFetcher userFetcher) *srv {
	return &srv{
		allConn:     make(map[int64][]userConn),
		userFetcher: userFetcher,
		pingTimeout: 10 * time.Second,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool { return true },
		},
	}
}

func (s *srv) Register(router *gin.RouterGroup) {
	router.GET("/ws", s.ws)
}
