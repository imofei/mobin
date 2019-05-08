package crons

import (
	"log"
	"github.com/imofei/gin-easy/comm/cron"
)

var (
	m *cron.Manager
)

func Start() {
	m = cron.NewManager()

	//m.Register(jobs.RewardListsMarkAlreadyGrant{})

	m.Start()
	log.Println("cron has started successfully")
}
