package remind_server

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
)

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
	fmt.Println("提醒成功，内容：", remindMsg)
	body.LastExecLibraryOffset = next
	body.LastExecTime = time.Now().Unix()
	body.CircleStart = true
	fmt.Printf("更改后的执行内容：%#v", body)
}

func (s *server) NeedRemind(tickerBody *TickerBody) bool {
	fmt.Println("上次提醒时间：", tickerBody.LastExecTime)
	fmt.Println("每隔几秒提醒：", tickerBody.CircleTime)
	fmt.Println("当前时间：", time.Now().Unix())
	if time.Now().Unix() >= tickerBody.LastExecTime+tickerBody.CircleTime {
		return true
	} else {
		return false
	}
}
