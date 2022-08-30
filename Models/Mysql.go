package Models

import (
	"fmt"
	"ginApi/Common/Logger"
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
	Logger.Println(sqlLog)
}

func init() {
	root := "root"
	password := "123456"
	dns := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local", root, password)

	newLogger := logger.New(
		Write{},
		//log.New(Tools.Logger, "", log.LstdFlags),
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
		panic(err)
	}
}
