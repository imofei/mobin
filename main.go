package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/imofei/gin-easy/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/imofei/gin-easy/comm/router"
	"github.com/imofei/gin-easy/app/config"
	"github.com/imofei/gin-easy/app/enums/codes"
	"github.com/imofei/gin-easy/app/http/controllers"
	"github.com/imofei/gin-easy/app/http/controllers/auth"
	"github.com/imofei/gin-easy/app/http/middlewares"
	"github.com/imofei/gin-easy/app/http/middlewares/jwt"
	"github.com/imofei/gin-easy/app/http/response"
)

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(config.AppMode())

	e := gin.New()
	if config.AppMode() == gin.DebugMode {
		e.Use(gin.Logger())
	}
	e.Use(middlewares.Recovery())
	e.Use(middlewares.Cors())

	e.NoRoute(func(c *gin.Context) {
		response.FailJSON(c, codes.NotFoundError, "no matched route")
	})

	// Gin原始的路由绑定方法
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"timestamp": time.Now().Unix(),
			"mode":      gin.Mode(),
		})
	})

	// html
	e.LoadHTMLGlob("templates/*")
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// 登录授权
	e.POST("/login/:code", api.GetAuth)

	r := router.NewSlimRouter(e)

	// Slim的路由绑定方法，基于控制器

	// 不需要授权
	noAuthRouter := r.RegisterGroup("/")
	{
		noAuthRouter.RegisterGroup("test").RegisterController(controllers.Test{}.NewController())
	}

	// 需要授权
	authRouter := r.RegisterGroup("/", jwt.JWT())
	{
		// 测试auth
		authRouter.RegisterGroup("test-auth").RegisterController(controllers.Test{}.NewController())
	}

	s := &http.Server{
		Addr:           config.AppListenAddress(),
		Handler:        http.TimeoutHandler(e, 30*time.Second, "request timeout"),
		ReadTimeout:    31 * time.Second,
		WriteTimeout:   31 * time.Second,
		MaxHeaderBytes: 1 << 19,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("listen fail", err)
	}
}
