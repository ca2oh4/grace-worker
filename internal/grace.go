package internal

import (
	"log"
	"os"
	"os/signal"
)

func Grace() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown...")
}
