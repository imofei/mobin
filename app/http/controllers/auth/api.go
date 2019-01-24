package api

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 测试授权
func CheckAuth(username, password string) bool {

	if username == "admin" && password == "admin123" {
		return true
	}

	return false
}



