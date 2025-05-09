package worker

import (
	"grace-worker/internal/worker/scheduler"
)

var s scheduler.Scheduler

func Run() error {
	s = scheduler.NewScheduler()
	return s.Run()
}

func Grace() {
	s.Stop()
}
