package kcs

import (
	"github.com/jinzhu/gorm"
	"github.com/imofei/gin-easy/app/components/db"
)

func MySQL() *gorm.DB {
	d, err := db.MySQL("kcs")
	if err != nil {
		panic(err)
	}
	return d
}
