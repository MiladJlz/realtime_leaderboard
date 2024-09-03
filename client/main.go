package main

import (
	"github.com/gorilla/websocket"
	"github.com/miladjlz/leaderboard/types"
	"log"
	"math/rand"
	"time"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

var sendInterval = time.Second * 5

func main() {
	users := generateUsers(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(users); i++ {
			score := generateUserScore()
			data := types.UserScore{UserID: users[i].ID, Score: score}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}

func generateUserScore() int {
	scores := []int{5, 10, 15, 20}
	randomIndex := rand.Intn(len(scores))
	randomValue := scores[randomIndex]
	return randomValue
}

func generateUsers(num int) []types.User {
	users := make([]types.User, num)
	for i := 0; i < num; i++ {
		users[i] = types.User{ID: int64(i) + 1, Username: "user" + string(rune(i+1))}
	}
	return users
}
