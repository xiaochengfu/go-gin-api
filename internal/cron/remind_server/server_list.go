package remind_server

import (
	"github.com/spf13/cast"
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
		libraryList, err = s.LibraryListByCategoryId(plan.CategoryId)
		if err != nil {
			return libraryList, err
		}
	} else {
		if plan.LibraryId != "" {
			libraryIds := strings.Split(plan.LibraryId, ",")
			var ids []int32
			for _, v := range libraryIds {
				ids = append(ids, cast.ToInt32(v))
			}
			libraryList, err = s.LibraryListByIds(ids)
		}
	}
	return libraryList, nil
}

func (s *server) LibraryListByCategoryId(categoryId int32) ([]*remind_library_repo.RemindLibrary, error) {
	query := remind_library_repo.NewQueryBuilder()
	list, err := query.WhereCategoryId(db_repo.EqualPredicate, categoryId).QueryAll(s.db.GetDbR())
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *server) LibraryListByIds(ids []int32) ([]*remind_library_repo.RemindLibrary, error) {
	query := remind_library_repo.NewQueryBuilder()
	list, err := query.WhereIdIn(ids).QueryAll(s.db.GetDbR())
	if err != nil {
		return nil, err
	}
	return list, nil
}
