package socket_server

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

func (s *server) OnMessage() {
	defer func() {
		s.OnClose()
	}()

	for {
		//接收消息
		_, message, err := s.socket.ReadMessage()
		if err != nil {
			s.logger.Error("socket on message error", zap.Error(err))
			break
		}
		err = processMsg(message)
		if err != nil {
			s.logger.Error("不合法的请求数据： " + string(message))
		}
		// 为了便于演示，仅输出到日志文件
		s.logger.Info("receive message: " + string(message))
	}
}

type Request struct {
	MsgId string `json:"msgId"` // 消息的唯一Id
	Cmd   string `json:"cmd"`   // 请求命令字
}

type LoginRequest struct {
	Request
	Data struct {
		Token string `json:"token"`
	}
}

func processMsg(message []byte) error {
	request := &Request{}
	err := json.Unmarshal(message, request)
	if err != nil {
		return err
	}
	switch request.Cmd {
	case "login":
		//验证用户，存储用户信息
		loginRequest := &LoginRequest{}
		err = json.Unmarshal(message, loginRequest)
		if err != nil {
			return err
		}
		userLogin(loginRequest)
	case "heart":
		//心跳检测
	}
	return nil
}

func userLogin(request *LoginRequest) {
	fmt.Println(request.Data.Token)
}
