package wshandlers

import "github.com/gorilla/websocket"

type message struct {
	text    string
	msgType int
	conn    *websocket.Conn
}

type userConn struct {
	conn *websocket.Conn
	ch   chan message
}

func (s *srv) addConn(conn userConn, userID int64) {
	s.Lock()
	defer s.Unlock()
	s.allConn[userID] = append(s.allConn[userID], conn)
}

func (s *srv) deleteConn(conn userConn, userID int64) {
	s.Lock()
	defer s.Unlock()
	for i, item := range s.allConn[userID] {
		if item.conn == conn.conn {
			s.allConn[userID] = append(s.allConn[userID][:i], s.allConn[userID][i+1:]...)
			return
		}
	}
}
