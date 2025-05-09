package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"grace-worker/internal/web/config"

	"grace-worker/pkg/database"
	"grace-worker/pkg/redis"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    config.Server.ToAddr(),
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
