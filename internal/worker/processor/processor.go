package processor

import (
	"context"
	"log"
	"sync"

	"grace-worker/pkg/redis"

	libredis "github.com/redis/go-redis/v9"
)

type Processor interface {
	Run() error
	Stop() error
}

type processor struct {
	done chan struct{}
	wg   sync.WaitGroup
}

func NewProcessor() Processor {
	return &processor{
		done: make(chan struct{}, 1),
	}
}

func (p *processor) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := redis.Client.XGroupCreateMkStream(ctx, "task-stream", "task-group", "$").Err()
	if err != nil {
		log.Printf("could not create consumer group: %v", err)
		return err
	}

	for {
		select {
		case <-p.done:
			log.Println("processor stopped")
			p.wg.Wait()
			return nil
		default:
			// 从 Stream 中读取消息
			entries, err := redis.Client.XReadGroup(ctx, &libredis.XReadGroupArgs{
				Group:    "task-group",
				Consumer: "consumer-1",
				Streams:  []string{"task-stream", ">"},
				Block:    0, // 阻塞直到有新消息
				Count:    1, // 每次获取 1 条消息
			}).Result()

			if err != nil {
				log.Printf("XReadGroup error: %v", err)
				continue
			}

			// 处理消息
			for _, entry := range entries {
				for _, msg := range entry.Messages {
					p.wg.Add(1)
					go p.handleTask(msg)
					// 确认消息
					redis.Client.XAck(ctx, "task-stream", "task-group", msg.ID)
				}
			}
		}
	}
}

func (p *processor) Stop() error {
	p.done <- struct{}{}
	return nil
}

func (p *processor) handleTask(msg libredis.XMessage) {
	defer p.wg.Done()

	log.Printf("handle task: %s", msg.ID)

	//	todo
}
