package wshandlers

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) UpdateTask(userID int64) {
	s.sendEvent(models.UpdatedTaskEventType, userID)
}

func (s *srv) CreateTask(userID int64) {
	s.sendEvent(models.CreatedTaskEventType, userID)
}

func (s *srv) sendEvent(eventType models.EventType, userID int64) {
	s.RLock()
	defer s.RUnlock()
	for _, conn := range s.allConn[userID] {
		go func() {
			select {
			case conn.ch <- message{
				text:    string(eventType),
				msgType: websocket.TextMessage,
				conn:    conn.conn,
			}:
			case <-time.After(time.Second * 2):
			}
		}()
	}
}

func (s *srv) send(ctx context.Context, conn userConn, userID int64) {
	defer s.deleteConn(conn, userID)
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-conn.ch:
			if msg.msgType == websocket.PingMessage {
				err := conn.conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					log.Println(err)
					return
				}
				continue
			}
			err := conn.conn.WriteMessage(websocket.TextMessage, []byte(msg.text))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
