package db

import (
	"gin-boilerplate/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitMySQL() {
	config := util.GetConfig()
	dsn := config.GetString("mysql.dsn")
	maxIdleConns := config.GetInt("mysql.maxIdleConns")
	maxOpenConns := config.GetInt("mysql.maxOpenConns")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(maxIdleConns)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(maxOpenConns)

	_db = db
}

func GetDB() *gorm.DB {
	return _db
}
