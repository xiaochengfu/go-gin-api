package remind_server

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_library_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/remind_plan_repo"
)

func (s *server) OnceRemind(libraries []*remind_library_repo.RemindLibrary) {
	remindMsg := ""
	for i := 0; i < len(libraries); i++ {
		remindMsg += libraries[i].Body
	}
	fmt.Println("提醒成功，内容：", remindMsg)
}

func (s *server) ClonePlan(id int32) error {
	update := map[string]interface{}{
		"status": id,
	}
	err := remind_plan_repo.NewQueryBuilder().WhereId(db_repo.EqualPredicate, id).Updates(s.db.GetDbW(), update)
	if err != nil {
		return err
	}
	return nil
}
