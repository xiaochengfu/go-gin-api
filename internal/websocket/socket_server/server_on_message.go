package socket_server

import (
	"encoding/json"
	"fmt"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
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
		s.logger.Info("receive message: " + string(message))
		err = s.processMsg(message)
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
}

type Request struct {
	MsgId string `json:"msg_id"` // 消息的唯一Id
	Cmd   string `json:"cmd"`    // 请求命令字
}

type LoginRequest struct {
	Request
	Data struct {
		Token    string `json:"token"`
		UserId   int64  `json:"user_id"`
		Username string `json:"username"`
	}
}

type LoginResponse struct {
	Request
	Response struct {
		UserId   int64  `json:"user_id"`
		Username string `json:"username"`
	} `json:"response"`
}

func (s *server) processMsg(message []byte) (err error) {
	request := &Request{}
	err = json.Unmarshal(message, request)
	if err != nil {
		return err
	}
	var jsonMsg []byte
	var userId int64
	switch request.Cmd {
	case "login":
		//验证用户，存储用户信息
		loginRequest := &LoginRequest{}
		err = json.Unmarshal(message, loginRequest)
		if err != nil {
			return err
		}
		var response *LoginResponse
		response, userId, err = userLogin(loginRequest)
		if err != nil {
			return errors.Wrap(err, "登录命令处理失败")
		}
		jsonMsg, err = json.Marshal(response)
		if err != nil {
			return err
		}
		clientManager.Users[userId] = s.socket
		fmt.Printf("Users:%v", clientManager.Users)
	}
	err = s.OnSend(userId, jsonMsg)
	if err != nil {
		return err
	}
	s.logger.Info("消息发送成功：" + string(jsonMsg))
	return
}

func userLogin(request *LoginRequest) (*LoginResponse, int64, error) {
	fmt.Println(request.Data.Token)
	//cfg := configs.Get().JWT
	//claims, err := token.New(cfg.Secret).JwtParse(request.Data.Token)
	//if err != nil {
	//	return nil, errors.Wrap(err,"token解析失败")
	//}
	userId := request.Data.UserId
	username := request.Data.Username
	loginResponse := &LoginResponse{
		Request: Request{
			Cmd:   request.Cmd,
			MsgId: request.MsgId,
		},
		Response: struct {
			UserId   int64  `json:"user_id"`
			Username string `json:"username"`
		}(struct {
			UserId   int64
			Username string
		}{UserId: userId, Username: username}),
	}
	return loginResponse, userId, nil
}
