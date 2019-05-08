package controllers

import (
	"fmt"
	"net/http"
	"time"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/imofei/gin-easy/comm/router"
	"github.com/imofei/gin-easy/app/http/response"
)

type Test struct {
	GetUser    func(c *gin.Context) `request:"GET /users"`
	CreateUser func(c *gin.Context) `request:"POST /users"`
	TestAuth   func(c *gin.Context) `request:"GET  /auth"`
	Test1      func(c *gin.Context) `request:"GET  /test1"`
}

func (t Test) NewController() router.Controller {
	t.GetUser = getUser
	t.CreateUser = createUser
	t.TestAuth = TestAuth
	t.Test1 = Test1
	return t
}

func getUser(c *gin.Context) {
	time.Sleep(40 * time.Second)
	panic(errors.New("test error"))
	c.String(http.StatusOK, "GET /users")
}

func createUser(c *gin.Context) {
	c.String(http.StatusOK, "POST /users")
}

func TestAuth(c *gin.Context) {

	data := make(map[string]interface{})
	data["lists"] = "test"
	data["total"] = "100"
	fmt.Println(c.Get("staff"))
	fmt.Println(data)

	response.SuccessJSON(c, data)

}

func Test1(c *gin.Context) {

	data := make(map[string]interface{})
	data["lists"] = "test"
	data["total"] = "100"
	fmt.Println(c.Get("staff"))
	fmt.Println(data)

	response.SuccessJSON(c, data)

}
