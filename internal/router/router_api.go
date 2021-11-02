package router

import (
	"github.com/xinliangnote/go-gin-api/internal/api/controller/admin_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/authorized_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/config_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/cron_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/menu_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/remind_plan_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/tool_handler"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func setApiRouter(r *resource) {

	// admin
	adminHandler := admin_handler.New(r.logger, r.db, r.cache)
	remindPlanHandler := remind_plan_handler.New(r.logger, r.db, r.cache)

	// 游客
	guest := r.mux.Group("/api/guest")
	{
		//提醒列表
		guest.GET("/remind/plan-list", remindPlanHandler.List())
	}

	// 需要签名验证，无需登录验证，无需 RBAC 权限验证
	login := r.mux.Group("/api", r.middles.Signature())
	{
		login.POST("/login", adminHandler.Login())
	}

	// 需要签名验证、登录验证，无需 RBAC 权限验证
	notRBAC := r.mux.Group("/api", core.WrapAuthHandler(r.middles.Token), r.middles.Signature())
	{
		notRBAC.POST("/admin/logout", adminHandler.Logout())
		notRBAC.PATCH("/admin/modify_password", adminHandler.ModifyPassword())
		notRBAC.GET("/admin/info", adminHandler.Detail())
		notRBAC.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())
	}

	// 需要签名验证、登录验证、RBAC 权限验证
	api := r.mux.Group("/api", core.WrapAuthHandler(r.middles.Token), r.middles.Signature(), r.middles.RBAC())
	{
		// authorized
		authorizedHandler := authorized_handler.New(r.logger, r.db, r.cache)
		api.POST("/authorized", authorizedHandler.Create())
		api.GET("/authorized", authorizedHandler.List())
		api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
		api.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())

		api.POST("/authorized_api", authorizedHandler.CreateAPI())
		api.GET("/authorized_api", authorizedHandler.ListAPI())
		api.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())

		api.POST("/admin", adminHandler.Create())
		api.GET("/admin", adminHandler.List())
		api.PATCH("/admin/used", adminHandler.UpdateUsed())
		api.PATCH("/admin/offline", adminHandler.Offline())
		api.PATCH("/admin/reset_password/:id", core.AliasForRecordMetrics("/api/admin/reset_password"), adminHandler.ResetPassword())
		api.DELETE("/admin/:id", core.AliasForRecordMetrics("/api/admin"), adminHandler.Delete())

		api.POST("/admin/menu", adminHandler.CreateAdminMenu())
		api.GET("/admin/menu/:id", core.AliasForRecordMetrics("/api/admin/menu"), adminHandler.ListAdminMenu())

		// menu
		menuHandler := menu_handler.New(r.logger, r.db, r.cache)
		api.POST("/menu", menuHandler.Create())
		api.GET("/menu", menuHandler.List())
		api.GET("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Detail())
		api.PATCH("/menu/used", menuHandler.UpdateUsed())
		api.PATCH("/menu/sort", menuHandler.UpdateSort())
		api.DELETE("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Delete())
		api.POST("/menu_action", menuHandler.CreateAction())
		api.GET("/menu_action", menuHandler.ListAction())
		api.DELETE("/menu_action/:id", core.AliasForRecordMetrics("/api/menu_action"), menuHandler.DeleteAction())

		// tool
		toolHandler := tool_handler.New(r.logger, r.db, r.cache)
		api.GET("/tool/hashids/encode/:id", core.AliasForRecordMetrics("/api/tool/hashids/encode"), toolHandler.HashIdsEncode())
		api.GET("/tool/hashids/decode/:id", core.AliasForRecordMetrics("/api/tool/hashids/decode"), toolHandler.HashIdsDecode())
		api.POST("/tool/cache/search", toolHandler.SearchCache())
		api.PATCH("/tool/cache/clear", toolHandler.ClearCache())
		api.GET("/tool/data/dbs", toolHandler.Dbs())
		api.POST("/tool/data/tables", toolHandler.Tables())
		api.POST("/tool/data/mysql", toolHandler.SearchMySQL())
		api.POST("/tool/send_message", toolHandler.SendMessage())

		// config
		configHandler := config_handler.New(r.logger, r.db, r.cache)
		api.PATCH("/config/email", configHandler.Email())

		// cron
		cronHandler := cron_handler.New(r.logger, r.db, r.cache, r.cronServer)
		api.POST("/cron", cronHandler.Create())
		api.GET("/cron", cronHandler.List())
		api.GET("/cron/:id", core.AliasForRecordMetrics("/api/cron/detail"), cronHandler.Detail())
		api.POST("/cron/:id", core.AliasForRecordMetrics("/api/cron/modify"), cronHandler.Modify())
		api.PATCH("/cron/used", cronHandler.UpdateUsed())
		api.PATCH("/cron/exec/:id", core.AliasForRecordMetrics("/api/cron/exec"), cronHandler.Execute())
	}
}
