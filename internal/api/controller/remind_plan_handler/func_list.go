package remind_plan_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type listRequest struct{}

type listResponse struct{}

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

	}
}
