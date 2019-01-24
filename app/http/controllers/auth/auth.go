package api

import (

	"github.com/gin-gonic/gin"

	"github.com/imofei/gin-easy/app/services"
	"github.com/imofei/gin-easy/app/http/response"
	"github.com/imofei/gin-easy/app/http/middlewares/jwt"
	"github.com/imofei/gin-easy/app/enums/codes"
)

type auth struct {
	Code string `form:"code" json:"code" binding:"required"`
}

func GetAuth(c *gin.Context) {
	tokenCode := c.Param("code")

	if err := len(tokenCode); err == 0 {
		response.FailJSON(c, codes.SystemError, "请求参数错误")
		return
	}

	authService := services.Auth{Code:tokenCode}
	err := authService.Check()

	if err != nil {
		response.FailJSON(c, codes.SystemError, err.Error())
		return
	}

	token, err := jwt.GenerateToken("015808531926232530")   // 测试用
	if err != nil {
		response.FailJSON(c, codes.SystemError, "Token生成失败")
		return
	}

	response.SuccessJSON(c, map[string]string{
		"token": token,
	})

}

