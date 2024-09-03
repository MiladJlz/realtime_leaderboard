package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"sync"
)

const (
	kafkaTopic = "game_score"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		store      StorerService
		ctx        = context.Background()
		connString = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			"localhost", 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASS"), "leaderboard", "disable",
		)
		wg = new(sync.WaitGroup)
	)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	store = NewStore(ctx, rdb, conn, wg)

	store = NewLogMiddleware(store)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, store)

	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()

}
