package remind

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/cron/remind_server"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/logger"
	"go.uber.org/zap"
)

type resource struct {
	mux    core.Mux
	logger *zap.Logger
	db     db.Repo
	cache  cache.Repo
}

type Server struct {
	Mux   core.Mux
	Db    db.Repo
	Cache cache.Repo
}

func NewServer(logger *zap.Logger) (*Server, error) {
	r := new(resource)
	r.logger = logger
	// 初始化 DB
	dbRepo, err := db.New()
	if err != nil {
		logger.Fatal("new db err", zap.Error(err))
	}
	r.db = dbRepo
	// 初始化 Cache
	cacheRepo, err := cache.New()
	if err != nil {
		logger.Fatal("new cache err", zap.Error(err))
	}
	r.cache = cacheRepo
	// gin引擎
	mux, err := core.New(logger,
		core.WithPanicNotify(notify.OnPanicNotify),
	)
	if err != nil {
		panic(err)
	}

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache
	return s, err
}

var (
	StartCmd = &cobra.Command{
		Use:     "remind",
		Short:   "提醒任务",
		Example: "remind start",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	// 初始化 cron logger
	cronLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.ProjectCronLogFile),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(cronLogger)
	s, _ := NewServer(cronLogger)

	remind := remind_server.New(s.Db)
	planList, err := remind.PlanList()
	if err != nil {
		return err
	}
	panic(planList)
	fmt.Println("remind ok")
	return nil
}
