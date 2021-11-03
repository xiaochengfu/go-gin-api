package remind_server

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_library_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
	"strings"
)

func (s *server) PlanList() (list []*remind_plan_repo.RemindPlan, err error) {
	query := remind_plan_repo.NewQueryBuilder()
	list, err = query.WhereStatus(db_repo.EqualPredicate, 1).
		QueryAll(s.db.GetDbR())
	if err != nil {
		return nil, err
	}
	return
}

func (s *server) LibraryListByPlan(plan *remind_plan_repo.RemindPlan) (libraryList []*remind_library_repo.RemindLibrary, err error) {
	if plan.CategoryId != 0 {
		//通过分组获取library
		libraryList, err = p.LibraryListByCategoryId(plan.CategoryId)
		if err != nil {
			return libraryList, errors.Wrap(err, "通过plan获取library失败")
		}
	} else {
		if plan.LibraryId != "" {
			libraryIds := strings.Split(plan.LibraryId, ",")
			libraryList, err = p.LibraryListByIds(libraryIds)
		}
	}
	return libraryList, nil
}
