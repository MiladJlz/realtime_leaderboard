package main

import (
	"github.com/miladjlz/leaderboard/types"
	"github.com/sirupsen/logrus"
	"time"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.UserScore) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"userID": data.UserID,
			"score":  data.Score,
			"took":   time.Since(start),
		}).Info("producing to kafka game_score topic")
	}(time.Now())
	return l.next.ProduceData(data)
}
