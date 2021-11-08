package remind

import (
	"fmt"
	"github.com/xinliangnote/go-gin-api/internal/router"
	"github.com/xinliangnote/go-gin-api/internal/websocket/socket_conn/system_message"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
	"github.com/xinliangnote/go-gin-api/internal/cron/remind_server"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/logger"

	"github.com/spf13/cobra"
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
			return nil
		},
	}
	TickerList = make(map[int32]*remind_server.TickerBody)
)

func Run(s *router.Server) error {
	timer := time.NewTicker(5 * time.Second)
	defer timer.Stop()
	for {
		conn, _ := system_message.GetConn()
		if conn == nil {
			continue
		}
		remind := remind_server.New(s.Db, conn)
		select {
		case <-timer.C:
			planList, err := remind.PlanList()
			onlyTime := time.Now().Format("15:04:00")
			if err != nil {
				return err
			}
			for _, v := range planList {
				fmt.Printf("%#v", v)
				planItem := v
				second := remind.ConvSecond(planItem.Time)
				//获取提醒信息
				libraries, err := remind.LibraryListByPlan(planItem)
				if err != nil {
					return err
				}
				if planItem.Type == remind_plan_repo.TypeSpecifyTime {
					if onlyTime == planItem.Time {
						//关闭任务
						err := remind.ClonePlan(planItem.Id)
						if err != nil {
							return err
						}
						remind.OnceRemind(libraries)
					}
				} else if planItem.Type == remind_plan_repo.TypeIntervalTime {
					if _, ok := TickerList[planItem.Id]; !ok {
						TickerList[planItem.Id] = &remind_server.TickerBody{
							CircleStart:           false,
							Circle:                int8(planItem.CircleType),
							LastExecTime:          time.Now().Unix(),
							CircleTime:            second,
							LibraryList:           libraries,
							LastExecLibraryOffset: 0,
						}
					}
					isNeed := remind.NeedRemind(TickerList[planItem.Id])
					if isNeed {
						remind.RepeatRemind(TickerList[planItem.Id])
					}
				}
			}
		}
	}

}

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

	timer := time.NewTicker(5 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			remind := remind_server.New(s.Db, nil)
			planList, err := remind.PlanList()
			onlyTime := time.Now().Format("15:04:00")
			if err != nil {
				return err
			}
			for _, v := range planList {
				fmt.Printf("%#v", v)
				planItem := v
				second := remind.ConvSecond(planItem.Time)
				//获取提醒信息
				libraries, err := remind.LibraryListByPlan(planItem)
				if err != nil {
					return err
				}
				if planItem.Type == remind_plan_repo.TypeSpecifyTime {
					if onlyTime == planItem.Time {
						//关闭任务
						err := remind.ClonePlan(planItem.Id)
						if err != nil {
							return err
						}
						remind.OnceRemind(libraries)
					}
				} else if planItem.Type == remind_plan_repo.TypeIntervalTime {
					if _, ok := TickerList[planItem.Id]; !ok {
						TickerList[planItem.Id] = &remind_server.TickerBody{
							CircleStart:           false,
							Circle:                int8(planItem.CircleType),
							LastExecTime:          time.Now().Unix(),
							CircleTime:            second,
							LibraryList:           libraries,
							LastExecLibraryOffset: 0,
						}
					}
					isNeed := remind.NeedRemind(TickerList[planItem.Id])
					if isNeed {
						remind.RepeatRemind(TickerList[planItem.Id])
					}
				}
			}
		}
	}

}
