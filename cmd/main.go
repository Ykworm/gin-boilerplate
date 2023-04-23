package main

import (
	"gin-boilerplate/common/db"
	"gin-boilerplate/model/entities"
	"gin-boilerplate/util"
	"github.com/prometheus/common/log"
)

func loadConfig(env string) *util.Config {
	config := util.GetConfig()
	config.InitConfig(env)
	return config
}

func main() {

	//config := loadConfig("dev")
	loadConfig("dev")
	db.InitMySQL()

	db := db.GetDB()
	var user entities.User
	db.Model(&user).Where("id = ?", "1267883854929068078").First(&user)
	log.Info(user)

}
