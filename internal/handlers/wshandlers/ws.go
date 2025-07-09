package wshandlers

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func (s *srv) ws(c *gin.Context) {
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	token := c.Query("token")
	user, err := s.userFetcher.GetUser(token)
	if err != nil {
		log.Println(err)
		return
	}
	uConn := userConn{
		conn: conn,
		ch:   make(chan message),
	}
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	s.addConn(uConn, user.ID)
	defer s.deleteConn(uConn, user.ID)
	go s.send(ctx, uConn, user.ID)
	go s.ping(ctx, uConn, user.ID)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, _, err := conn.ReadMessage()
			if err == nil {
				continue
			}
			log.Println("ws disconnected:", err)
			return
		}
	}
}
