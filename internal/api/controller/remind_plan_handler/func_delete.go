package remind_plan_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 删除计划
// @Summary 删除计划
// @Description 删除计划
// @Tags API.remind
// @Accept json
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/remind/delete [post]
func (h *handler) Delete() core.HandlerFunc {
	return func(c core.Context) {

	}
}
