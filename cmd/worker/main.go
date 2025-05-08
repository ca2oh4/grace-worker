package main

import (
	"log"

	"grace-worker/pkg/runtime"
)

func main() {
	runtime.Grace()
	log.Println("Shutdown Worker ...")
}
