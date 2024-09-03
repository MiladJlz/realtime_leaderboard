package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%d dbname=%s sslmode=%s",
		"localhost", 5432, "postgres", 159753, "leaderboard", "disable",
	)
	conn, err := sql.Open("postgres", connString)

	for i := 0; i < 20; i++ {
		_, err = conn.ExecContext(context.Background(), "INSERT INTO users (username) VALUES ($1)", "user"+fmt.Sprint(i))

		if err != nil {
			log.Fatal(err)
		}
	}

}
