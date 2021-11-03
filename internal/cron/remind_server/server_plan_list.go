package remind_server

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
)

func (s *server) PlanList() (list []*remind_plan_repo.RemindPlan, err error) {
	query := remind_plan_repo.NewQueryBuilder()
	list, err = query.WhereStatus(db_repo.EqualPredicate, 1).
		QueryAll(s.db.GetDbR().WithContext(s.ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
