package remind_server

import (
	"encoding/json"
	"fmt"
	"github.com/xinliangnote/go-gin-api/internal/websocket/socket_server"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
)

type RemindResponse struct {
	socket_server.Request
	Response struct {
		Body string `json:"body"`
		Msg  string `json:"msg"`
	} `json:"response"`
}

func (s *server) RepeatRemind(body *TickerBody) {
	var next int
	if body.Circle == remind_plan_repo.RemindSort {
		//顺序
		next = body.LastExecLibraryOffset + 1
	} else if body.Circle == remind_plan_repo.RemindRand {
		//随机
		next = rand.Intn(len(body.LibraryList))
		if next == body.LastExecLibraryOffset {
			next += 1
		}
	}
	if next >= len(body.LibraryList) {
		next = 0
	}
	library := body.LibraryList[next]
	remindMsg := library.Body
	FeiShuMsg := &FeiShuMsg{
		MsgType: "interactive",
		Card: struct {
			Config struct {
				WideScreenMode bool `json:"wide_screen_mode"`
			} `json:"config"`
			Header struct {
				Template string `json:"template"`
				Title    struct {
					Content string `json:"content"`
					Tag     string `json:"tag"`
				} `json:"title"`
			} `json:"header"`
			I18NElements struct {
				ZhCn []struct {
					Tag  string `json:"tag"`
					Text struct {
						Content string `json:"content"`
						Tag     string `json:"tag"`
					} `json:"text"`
				} `json:"zh_cn"`
			} `json:"i18n_elements"`
		}{Config: struct {
			WideScreenMode bool `json:"wide_screen_mode"`
		}{true}, Header: struct {
			Template string `json:"template"`
			Title    struct {
				Content string `json:"content"`
				Tag     string `json:"tag"`
			} `json:"title"`
		}{Template: "orange", Title: struct {
			Content string `json:"content"`
			Tag     string `json:"tag"`
		}{remindMsg, "plain_text"}}, I18NElements: struct {
			ZhCn []struct {
				Tag  string `json:"tag"`
				Text struct {
					Content string `json:"content"`
					Tag     string `json:"tag"`
				} `json:"text"`
			} `json:"zh_cn"`
		}{ZhCn: []struct {
			Tag  string `json:"tag"`
			Text struct {
				Content string `json:"content"`
				Tag     string `json:"tag"`
			} `json:"text"`
		}{struct {
			Tag  string `json:"tag"`
			Text struct {
				Content string `json:"content"`
				Tag     string `json:"tag"`
			} `json:"text"`
		}{Tag: "div", Text: struct {
			Content string `json:"content"`
			Tag     string `json:"tag"`
		}{"提醒次数：1次 \n提醒类型：重复提醒", "lark_md"}}}}},
	}
	jsonResponse, _ := json.Marshal(FeiShuMsg)
	fmt.Println(string(jsonResponse))
	resp, err := http.Post("https://open.feishu.cn/open-apis/bot/v2/hook/48ad45fc-c177-4899-a7a4-f78365ac0e77",
		"application/json",
		strings.NewReader(string(jsonResponse)))
	if err != nil {
		fmt.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(string(body))
			}
		}
	}(resp.Body)
	//err := s.wsConnect.OnSend(int64(library.UserId), jsonResponse)
	//if err != nil {
	//	fmt.Println("推送失败")
	//}
	fmt.Println("提醒成功，内容：", remindMsg)
	body.LastExecLibraryOffset = next
	body.LastExecTime = time.Now().Unix()
	body.CircleStart = true
	//fmt.Printf("更改后的执行内容：%#v", body)
}

func (s *server) NeedRemind(tickerBody *TickerBody) bool {
	//fmt.Println("上次提醒时间：", tickerBody.LastExecTime)
	fmt.Println("每隔几秒提醒：", tickerBody.CircleTime)
	//fmt.Println("当前时间：", time.Now().Unix())
	if time.Now().Unix() >= tickerBody.LastExecTime+tickerBody.CircleTime {
		return true
	} else {
		return false
	}
}

type FeiShuMsg struct {
	MsgType string `json:"msg_type"`
	Card    struct {
		Config struct {
			WideScreenMode bool `json:"wide_screen_mode"`
		} `json:"config"`
		Header struct {
			Template string `json:"template"`
			Title    struct {
				Content string `json:"content"`
				Tag     string `json:"tag"`
			} `json:"title"`
		} `json:"header"`
		I18NElements struct {
			ZhCn []struct {
				Tag  string `json:"tag"`
				Text struct {
					Content string `json:"content"`
					Tag     string `json:"tag"`
				} `json:"text"`
			} `json:"zh_cn"`
		} `json:"i18n_elements"`
	} `json:"card"`
}
