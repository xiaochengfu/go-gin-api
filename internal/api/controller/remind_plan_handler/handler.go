package remind_plan_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 新增计划
	// @Tags API.remind
	// @Router /api/remind/create [post]
	Create() core.HandlerFunc

	// Delete 删除计划
	// @Tags API.remind
	// @Router /api/remind/delete [post]
	Delete() core.HandlerFunc

	// Update 修改计划
	// @Tags API.remind
	// @Router /api/remind/update [post]
	Update() core.HandlerFunc

	// List 计划列表
	// @Tags API.remind
	// @Router /api/remind [get]
	List() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	cache  cache.Repo
}

func (h *handler) i() {}
