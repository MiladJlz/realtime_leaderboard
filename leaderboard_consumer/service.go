package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/miladjlz/leaderboard/types"
	"log"
	"net/http"
)

type DataSender struct {
	msgch chan []types.LeaderBoard
	conn  *websocket.Conn
	kc    *KafkaConsumer
}

func NewDataSender() (*DataSender, error) {
	var (
		c          *KafkaConsumer
		kafkaTopic = "leaderboard"
		msgch      = make(chan []types.LeaderBoard, 128)
	)

	c, err := NewKafkaConsumer(kafkaTopic, msgch)
	if err != nil {
		return nil, err

	}
	return &DataSender{kc: c,
		msgch: msgch,
	}, nil
}

func (ds *DataSender) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	},
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	ds.conn = conn
	go ds.wsSendLoop()
}

func (ds *DataSender) wsSendLoop() {
	fmt.Println("New client connected !")
	go ds.kc.Start()
	for {
		select {
		case value := <-ds.msgch:
			if err := ds.conn.WriteJSON(value); err != nil {
				log.Fatal(err)
			}
		}
	}
}
