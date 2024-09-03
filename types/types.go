package types

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UserScore struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"userID"`
	Score  int   `json:"score"`
}
type LeaderBoard struct {
	Score  float64     `json:"score"`
	Member interface{} `json:"member"`
}
