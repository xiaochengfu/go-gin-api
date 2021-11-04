package cron_server

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/cron_task_repo"

	"github.com/jakecoffman/cron"
)

func (s *server) AddJob(task *cron_task_repo.CronTask) cron.FuncJob {
	return func() {
		s.taskCount.Add()
		defer s.taskCount.Done()

		// 将 task 信息写入到 Kafka Topic 中，任务执行器订阅 Topic 如果为符合条件的任务并进行执行，反之不执行
		// 为了便于演示，不写入到 Kafka 中，仅记录日志
		dir, _ := os.Getwd()
		projectPath := strings.Replace(dir, "\\", "/", -1)
		bash := projectPath + "/scripts/bash.sh"
		shellPath := fmt.Sprintf("%s %s", bash, task.Command)
		command := new(exec.Cmd)

		// runtime.GOOS = linux or darwin
		command = exec.Command("/bin/bash", "-c", shellPath)

		var stderr bytes.Buffer
		command.Stderr = &stderr

		output, _ := command.Output()
		msg := fmt.Sprintf("执行任务：(%d)%s [%s] [执行命令:%s] [result:%s]", task.Id, task.Name, task.Spec, shellPath, string(output))
		s.logger.Info(msg)
	}
}
