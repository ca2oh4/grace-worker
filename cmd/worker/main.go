package main

import (
	"log"

	"grace-worker/internal"
)

func main() {
	internal.Grace()
	log.Println("Shutdown Worker ...")
}
