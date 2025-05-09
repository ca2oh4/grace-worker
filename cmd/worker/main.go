package main

import (
	"log"

	"grace-worker/internal/worker"
	"grace-worker/internal/worker/config"

	"grace-worker/pkg/database"
	"grace-worker/pkg/redis"
	"grace-worker/pkg/runtime"
)

func main() {
	if err := config.Setup(); err != nil {
		log.Fatalln(err)
	}
	if err := database.Setup(config.Database.ToDatabaseOption()); err != nil {
		log.Fatalln(err)
	}
	if err := redis.Setup(config.Redis.ToRedisOptions()); err != nil {
		log.Fatalln(err)
	}

	defer func() {
		re := recover()
		if re != nil {
			log.Println("worker panic:", re)
		}
		worker.Grace()
		log.Println("Shutdown Worker ...")
	}()

	go func() {
		if err := worker.Run(); err != nil {
			log.Panicln(err)
		}
	}()

	runtime.Grace()
}
