package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/miladjlz/leaderboard/types"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

type StorerService interface {
	InsertScore(types.UserScore)
}
type Store struct {
	ctx context.Context
	ps  *sql.DB
	rs  *redis.Client
	wg  *sync.WaitGroup
}

func NewStore(ctx context.Context, rconn *redis.Client, pconn *sql.DB, wg *sync.WaitGroup) *Store {
	return &Store{ctx, pconn, rconn, wg}
}

func (s *Store) InsertScore(data types.UserScore) {

	s.wg.Add(2)
	go s.InsertToPostgres(data, s.ctx, s.wg)
	go s.InsertToRedis(data, s.ctx, s.wg)
	s.wg.Wait()
}

func (s *Store) InsertToPostgres(data types.UserScore, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := s.ps.ExecContext(ctx, "INSERT INTO users_scores (score, user_id) VALUES ($1, $2)", data.Score, data.UserID)
	if err != nil {
		log.Fatal(err)
	}
}
func (s *Store) InsertToRedis(data types.UserScore, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	err := s.rs.ZIncrBy(ctx, "game_score", float64(data.Score), fmt.Sprint(data.UserID)).Err()
	if err != nil {
		log.Fatal(err)
	}
}
