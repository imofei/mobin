gin-slim-router
======
A slim router for gin framework

## Usage

1. Create controller
```Go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"http://gitlab.kucoin.net/golang/router"
)

type User struct {
	// HttpMethod: Any/GET/POST/DELETE/PATCH/PUT/OPTIONS/HEAD
	// Format: Action func(*gin.Context) `request:"{HttpMethod} /xxxpath"`
	Create func(*gin.Context) `request:"post /create"`
	List   func(*gin.Context) `request:"get /list"`
}

func (u User) NewController() router.Controller {
	u.Create = create
	u.List = getList
	return u
}

func create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "create user")
}

func getList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get user list")
}
```

2. Register route
```Go
package main

import (
	"github.com/gin-gonic/gin"
	"http://gitlab.kucoin.net/golang/examples/router/controllers"
	"http://gitlab.kucoin.net/golang/router"
)

func main() {
	ng := gin.New()
	r := router.NewSlimRouter(ng)
	// r.Use(middlewares.Global()) // Global middleware

	// Register route: POST http://127.0.0.1:5200/xxx
	// Register route: GET http://127.0.0.1:5200/yyy
	// Register route: ANY http://127.0.0.1:5200/any
	// r.RegisterController(controllers.Test{}.NewController())

	// Register route: POST http://127.0.0.1:5200/user/create
	// Register route: GET http://127.0.0.1:5200/user/list
	r.RegisterGroup("user" /*, middlewares.Auth()*/).RegisterController(controllers.User{}.NewController())
	ng.Run(":5200")
}
```

## License

[MIT](https://github.com/hhxsv5/gin-slim-router/blob/master/LICENSE)
