package wshandlers

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func (s *srv) ping(ctx context.Context, uConn userConn, userID int64) {
	defer s.deleteConn(uConn, userID)
	resp := make(chan struct{}, 1)
	uConn.conn.SetPongHandler(func(_ string) error {
		select {
		case resp <- struct{}{}:
		case <-ctx.Done():
		}
		return nil
	})
	for {
		select {
		case <-time.After(s.pingTimeout / 2):
		case <-ctx.Done():
			return
		}
		select {
		case uConn.ch <- message{
			msgType: websocket.PingMessage,
			conn:    uConn.conn,
		}:
		case <-ctx.Done():
			return
		}
		select {
		case <-resp:
		case <-time.After(s.pingTimeout):
			log.Println("responce timeout")
			return
		case <-ctx.Done():
			return
		}
	}
}
