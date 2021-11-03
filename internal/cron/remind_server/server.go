package remind_server

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Server = (*server)(nil)

type Server interface {
	i()
	RepeatRemind()
	OnceRemind()
	PlanList() (list []*remind_plan_repo.RemindPlan, err error)
	//LibraryListByPlan()
}

type server struct {
	db  db.Repo
	ctx core.Context
}

func New(db db.Repo) *server {
	return &server{
		db: db,
	}
}

//func (s *server) LibraryListByPlan() {
//	panic("implement me")
//}

func (s *server) i() {}
