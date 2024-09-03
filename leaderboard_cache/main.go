package main

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/miladjlz/leaderboard/types"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func main() {
	var (
		ctx = context.Background()
	)
	dr, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	rs := NewRedisStore(rdb.Conn(), ctx)

	for {
		res, err := rs.GetTopScores()
		if err != nil {
			log.Fatal(err)
		}
		err = dr.prod.ProduceData(res)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 5)
	}

}

type DataReceiver struct {
	conn *websocket.Conn
	prod DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		p          DataProducer
		err        error
		kafkaTopic = "leaderboard"
	)
	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}
	p = NewLogMiddleware(p)
	return &DataReceiver{
		prod: p,
	}, nil
}

func (dr *DataReceiver) produceData(users []types.LeaderBoard) error {
	return dr.prod.ProduceData(users)
}
