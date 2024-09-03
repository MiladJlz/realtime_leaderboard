package main

import (
	"github.com/miladjlz/leaderboard/types"
	"github.com/sirupsen/logrus"
	"time"
)

type LogMiddleware struct {
	next StorerService
}

func NewLogMiddleware(next StorerService) StorerService {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) InsertScore(data types.UserScore) {

	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"userID": data.UserID,
			"score":  data.Score,
			"took":   time.Since(start),
		}).Info("consuming game_score topic and storing in db's")
	}(time.Now())
	m.next.InsertScore(data)
}
