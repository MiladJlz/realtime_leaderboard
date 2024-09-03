package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		"localhost", 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASS"), "leaderboard", "disable",
	)
	conn, err := sql.Open("postgres", connString)

	for i := 0; i < 20; i++ {
		_, err = conn.ExecContext(context.Background(), "INSERT INTO users (username) VALUES ($1)", "user"+fmt.Sprint(i))

		if err != nil {
			log.Fatal(err)
		}
	}

}
