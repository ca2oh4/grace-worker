package worker

import (
	"grace-worker/internal/worker/scheduler"
)

var s scheduler.Scheduler

func Run() error {
	s = scheduler.NewScheduler()

	// 获取执行权限
	acquired, err := s.Acquire()
	if err != nil {
		return err
	}
	if !acquired {
		return nil
	}

	// 调度任务执行
	return s.Run()
}

func Grace() error {
	return s.Stop()
}
