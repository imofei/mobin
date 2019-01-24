package services

type Auth struct {
	Code string
}

// 检查钉钉code, 首次打开注册并登录
func (a *Auth) Check() (error) {

	//client:=dingtalk.NewClient()
	//info, err := client.UserInfoByCode(a.Code)
	return nil
}
