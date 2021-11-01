package remind_plan_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改计划
// @Summary 修改计划
// @Description 修改计划
// @Tags API.remind
// @Accept json
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} updateResponse
// @Failure 400 {object} code.Failure
// @Router /api/remind/update [post]
func (h *handler) Update() core.HandlerFunc {
	return func(c core.Context) {

	}
}
