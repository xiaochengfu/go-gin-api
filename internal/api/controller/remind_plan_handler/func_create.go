package remind_plan_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type createRequest struct{}

type createResponse struct{}

// Create 新增计划
// @Summary 新增计划
// @Description 新增计划
// @Tags API.remind
// @Accept json
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/remind/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {

	}
}
