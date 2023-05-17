package models

import (
	"fmt"
	"ginApi/common/config"
	"ginApi/common/enum"
	"ginApi/common/myLogger"
	"ginApi/common/response"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Mysql struct{}

var (
	DB  *gorm.DB
	err error
)

type Write struct {
}

/**
filePath := values[0] // 文件路径
time := values[1]     // 执行时间
sql := values[3]      // sql语句
*/

func (w Write) Printf(format string, values ...interface{}) {
	sqlLog := fmt.Sprintf("mysql[%s] time[%v] filePath[%s]", values[3], values[1], values[0])
	myLogger.Println(sqlLog)
}

func init() {
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		config.Viper.Get("mysql.user"),
		config.Viper.Get("mysql.password"),
		config.Viper.Get("mysql.host"),
		config.Viper.Get("mysql.port"),
		config.Viper.Get("mysql.db"),
		config.Viper.Get("mysql.chaSet"))

	newLogger := logger.New(
		Write{},
		//log.New(tools.myLogger, "", log.LstdFlags),
		logger.Config{
			SlowThreshold:             800 * time.Millisecond,
			LogLevel:                  logger.Info,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
		},
	)

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println("mysql conn err:", err)
		panic(response.Response{Code: enum.CodeSystemError, Msg: "mysql conn err:"})
	}
}
