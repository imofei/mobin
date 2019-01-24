package redis

type Passport struct {
	UserId  uint64       `json:"userId"`
	Session string       `json:"session"`
	Data    PassportData `json:"data"`
}

type PassportData struct {
	UserId         uint64 `redis:"user_id" json:"user_id"`
	LoginTime      int64  `redis:"login_time" json:"login_time"`
	Platform       string `redis:"platform" json:"platform"`
	LastActiveTime int64  `redis:"last_active_time" json:"last_active_time"`
}
