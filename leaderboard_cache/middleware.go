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

func (l *LogMiddleware) ProduceData(data []types.LeaderBoard) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"users": data,
			"took":  time.Since(start),
		}).Info("caching to leaderboard redis")
	}(time.Now())

	return l.next.ProduceData(data)
}
