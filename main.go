package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/online.scheduling-worker/messenger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	messenger.Work_RemoveSchedulesFrom(context.Background())
}
