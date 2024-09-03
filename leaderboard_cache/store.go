package main

import (
	"context"
	"github.com/miladjlz/leaderboard/types"
	"github.com/redis/go-redis/v9"
	"log"
)

type CacheService interface {
	GetTopScores() ([]types.LeaderBoard, error)
	CacheTopScores() error
}

type RedisStore struct {
	conn *redis.Conn
	ctx  context.Context
}

func NewRedisStore(conn *redis.Conn, ctx context.Context) CacheService {
	return &RedisStore{conn: conn, ctx: ctx}
}
func (s *RedisStore) GetTopScores() ([]types.LeaderBoard, error) {
	var leaderboard []types.LeaderBoard
	exists, err := s.conn.Exists(s.ctx, "leaderboard").Result()
	if exists == 1 {
		err := s.conn.ZRemRangeByScore(s.ctx, "leaderboard", "-inf", "+inf").Err()
		if err != nil {
			panic(err)
		}
	}
	err = s.CacheTopScores()
	if err != nil {
		return nil, err
	}
	members, err := s.conn.ZRevRangeWithScores(s.ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		return nil, err

	}
	for _, member := range members {
		leaderboard = append(leaderboard, types.LeaderBoard{Score: member.Score, Member: member.Member})
	}
	return leaderboard, nil

}

func (s *RedisStore) CacheTopScores() error {

	topMembers, err := s.conn.ZRevRangeWithScores(s.ctx, "game_score", 0, 9).Result()
	if err != nil {
		log.Fatal(err)
	}

	for _, member := range topMembers {
		_, err = s.conn.ZAdd(s.ctx, "leaderboard", redis.Z{Score: member.Score, Member: member.Member}).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
