package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

const (
	kafkaTopic = "game_score"
)

func main() {
	var (
		store      StorerService
		ctx        = context.Background()
		connString = fmt.Sprintf(
			"host=%s port=%d user=%s password=%d dbname=%s sslmode=%s",
			"localhost", 5432, "postgres", 159753, "leaderboard", "disable",
		)
		wg = new(sync.WaitGroup)
	)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
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
