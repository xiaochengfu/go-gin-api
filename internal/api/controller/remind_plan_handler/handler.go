package remind_plan_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/api/service/remind_plan_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
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
	logger            *zap.Logger
	cache             cache.Repo
	remindPlanService remind_plan_service.Service
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:            logger,
		cache:             cache,
		remindPlanService: remind_plan_service.New(db),
	}
}

func (h *handler) i() {}
