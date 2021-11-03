package remind_plan_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/service/remind_plan_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/spf13/cast"
)

type listRequest struct {
	Page     int `form:"page"`      // 第几页
	PageSize int `form:"page_size"` // 每页显示条数
}

type listData struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	UserId     int    `json:"user_id"`
	CategoryId int    `json:"category_id"` //组id
	LibraryId  string `json:"library_id"`
	Time       string `json:"time"`        //时间时分秒
	Status     int8   `json:"status"`      //状态 1-开启 2-关闭 3-删除
	Type       int8   `json:"type"`        //提醒策略类型 1-指定时间 2-每x几分提醒一次
	CircleType int8   `json:"circle_type"` //多个提醒项时，周期内 1-随机 2-顺序
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

// List 计划列表
// @Summary 计划列表
// @Description 计划列表
// @Tags API.remind
// @Accept json
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/remind [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}

		searchData := new(remind_plan_service.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		resultListData, err := h.remindPlanService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.RemindPlanListError,
				code.Text(code.RemindPlanListError)).WithErr(err),
			)
			return
		}
		resCountData, err := h.remindPlanService.PageListCount(c, searchData) //todo 分页查询条件要写两份了
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.RemindPlanListError,
				code.Text(code.RemindPlanListError)).WithErr(err),
			)
			return
		}
		res.Pagination.Total = cast.ToInt(resCountData)
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(resultListData))
		for k, v := range resultListData {
			data := listData{
				Id:         cast.ToInt(v.Id),
				Name:       v.Name,
				UserId:     cast.ToInt(v.UserId),
				CategoryId: cast.ToInt(v.CategoryId),
				LibraryId:  v.LibraryId,
				Time:       v.Time,
				Status:     cast.ToInt8(v.Status),
				Type:       cast.ToInt8(v.Type),
				CircleType: cast.ToInt8(v.CircleType),
			}
			res.List[k] = data
		}
		c.Payload(res)
	}
}
