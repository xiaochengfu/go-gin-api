package remind_plan_service

import (
	"fmt"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchData struct {
	Page     int // 第几页
	PageSize int // 每页显示条数
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*remind_plan_repo.RemindPlan, err error) {
	page := searchData.Page
	if page == 0 {
		page = 1
	}
	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	query := remind_plan_repo.NewQueryBuilder()
	listData, err = query.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", listData)
	return
}
