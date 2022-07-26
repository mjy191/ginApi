package Models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Mysql struct{}

var (
	DB  *gorm.DB
	err error
)

func init() {
	root := "root"
	password := "123456"
	dns := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local", root, password)
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("mysql conn err:", err)
		panic(err)
	}
}
