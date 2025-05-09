package scheduler

import (
	"log"
	"time"

	"grace-worker/internal/worker/processor"
	"grace-worker/pkg/redis"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

type Scheduler interface {
	Acquire() (bool, error)
	Run() error
	Stop() error
}

type scheduler struct {
	mutex *redsync.Mutex

	acquired bool
	done     chan struct{}

	p processor.Processor
}

func NewScheduler() Scheduler {
	return &scheduler{
		done: make(chan struct{}, 1),
	}
}

func (s *scheduler) Acquire() (bool, error) {
	const (
		mutexName   = "worker-run-mutex"
		mutexExpiry = 10 * time.Second
	)

	pool := goredis.NewPool(redis.Client)
	r := redsync.New(pool)

	s.mutex = r.NewMutex(mutexName, redsync.WithExpiry(mutexExpiry))
	if err := s.mutex.Lock(); err != nil {
		log.Printf("lock fail")
		return false, nil
	}
	s.acquired = true

	go func() {
		ticker := time.NewTicker(mutexExpiry / 2)
		defer func() {
			ticker.Stop()

			// 释放锁
			s.acquired = false
			if ok, err := s.mutex.Unlock(); err != nil || !ok {
				log.Printf("unlock fail")
			}
		}()

		for {
			select {
			case <-ticker.C:
				ok, err := s.mutex.Extend()
				if err != nil {
					return
				}
				if !ok {
					return
				}
			case <-s.done:
				return
			}
		}
	}()

	return true, nil
}

func (s *scheduler) Run() error {
	if !s.acquired {
		return nil
	}

	s.p = processor.NewProcessor()
	if err := s.p.Run(); err != nil {
		return err
	}

	return nil
}

func (s *scheduler) Stop() error {
	if !s.acquired {
		return nil
	}

	// 先停止执行，再释放锁
	if s.p != nil {
		err := s.p.Stop()
		if err != nil {
			log.Printf("stop fail err: %+v", err)
		}
	}

	s.done <- struct{}{}

	return nil
}
