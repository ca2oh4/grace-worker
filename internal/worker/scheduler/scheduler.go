package scheduler

type Scheduler interface {
	Run() error
	Stop() error
}

func NewScheduler() Scheduler {
	return &scheduler{}
}

type scheduler struct{}

func (s *scheduler) Run() error {
	panic("todo")
}

func (s *scheduler) Stop() error {
	panic("todo")
}
