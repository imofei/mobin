package crons

import (
	"log"
	"gitlab.kucoin.net/golang/cron"
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
