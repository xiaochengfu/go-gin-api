package remind_plan_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	PageList(ctx core.Context, searchData *SearchData) (listData []*remind_plan_repo.RemindPlan, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
}

type service struct {
	db db.Repo
}

func (s service) i() {}

func New(db db.Repo) *service {
	return &service{
		db: db,
	}
}
