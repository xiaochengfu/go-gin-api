package socket_server

import "github.com/gorilla/websocket"

var clientManager = NewClientManage()

type clientManage struct {
	Users map[int64]*websocket.Conn
}

func NewClientManage() *clientManage {
	return &clientManage{
		Users: make(map[int64]*websocket.Conn),
	}
}
