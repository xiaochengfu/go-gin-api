package socket_server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"strconv"
)

func (s *server) OnSend(userId int64, message []byte) error {
	//err := s.socket.WriteMessage(websocket.TextMessage, message)
	fmt.Println("给用户" + strconv.Itoa(int(userId)) + "推送")
	fmt.Println("users列表为:", s.Users)
	if client, ok := s.Users[userId]; ok {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			s.OnClose()
			s.logger.Error("socket on send error", zap.Error(err))
		}
		return err
	} else {
		fmt.Println("用户" + strconv.Itoa(int(userId)) + "不存在")
		return nil
	}
}
