package bootstrap

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/imofei/gin-easy/app/config"
	"github.com/imofei/gin-easy/app/cron"
	"github.com/imofei/gin-easy/app/components/log"
	_ "github.com/imofei/gin-easy/app/components/db"
	_ "github.com/imofei/gin-easy/app/components/redis"
)

func init() {
	go writePIdFile(config.BasePath + "/runtime/pid")

	if err := log.InitDailyLog("gin", config.BasePath+"/runtime/logs"); err != nil {
		panic(err)
	}

	if config.AppCron() {
		crons.Start()
	}
}

// 记录PID到文件，Reload时需要
func writePIdFile(path string) {
	p := strconv.Itoa(os.Getpid())
	b := bytes.NewBufferString(p).Bytes()
	err := ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		panic(errors.New(fmt.Sprintf("create pidfile fail: %s", err)))
	}
}
