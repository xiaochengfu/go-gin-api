package remind_server

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
)

func (s *server) OnceRemind(item *remind_plan_repo.RemindPlan) {

}

func (s *server) ClonePlan(id int32) error {
	update := map[string]interface{}{
		"status": id,
	}
	err := remind_plan_repo.NewQueryBuilder().WhereId(db_repo.EqualPredicate, id).Updates(s.db.GetDbW(), update)
	if err != nil {
		return err
	}
	return nil
}
