package remind_server

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_library_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/websocket/socket_server"
)

var _ Server = (*server)(nil)

type Server interface {
	i()
	RepeatRemind(body *TickerBody)
	OnceRemind(libraries []*remind_library_repo.RemindLibrary)
	PlanList() (list []*remind_plan_repo.RemindPlan, err error)
	ConvSecond(t string) int64
	LibraryListByPlan(plan *remind_plan_repo.RemindPlan) (libraryList []*remind_library_repo.RemindLibrary, err error)
	LibraryListByCategoryId(categoryId int32) ([]*remind_library_repo.RemindLibrary, error)
	NeedRemind(tickerBody *TickerBody) bool
}

type TickerBody struct {
	CircleStart           bool                                 //循环是否开始
	Circle                int8                                 //周期性循环类型
	LastExecTime          int64                                //上次执行时间
	CircleTime            int64                                //多少秒执行一次
	LibraryList           []*remind_library_repo.RemindLibrary //提醒库列表
	LastExecLibraryOffset int                                  //上次执行提醒库的位置
}

type server struct {
	db        db.Repo
	wsConnect socket_server.Server
}

func New(db db.Repo, wsConnect socket_server.Server) *server {
	return &server{
		db:        db,
		wsConnect: wsConnect,
	}
}

func (s *server) i() {}
