package myLogger

import (
	"fmt"
	"ginApi/common/config"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

var LogId string
var mylog *log.Logger
var env bool = false

func init() {
	mylog = log.New(LogWrite, "", log.LstdFlags)
	if config.Viper.Get("env") == "prod" {
		env = true
	}
}

/*
GenerateLogId 生产日志ID
*/
func GenerateLogId() string {
	timeStr := time.Now().Format("20060102150405")
	// 生成UUID并去除横线
	uuidStr := strings.Replace(uuid.New().String(), "-", "", -1)

	// 截取UUID前几位（例如8位）+ 时间戳
	shortUuid := uuidStr[:13]

	// 组合格式：时间戳+短UUID
	LogId = fmt.Sprintf("%s%s", timeStr, shortUuid)

	return LogId
}

func Printf(msg string, v ...any) {
	if v != nil {
		msg = fmt.Sprintf(msg, v)
	}
	if !env {
		fmt.Println("logid["+LogId+"] ", msg)
	}
	mylog.Println("logid["+LogId+"] ", msg)
}

func Println(msg string) {
	if !env {
		fmt.Println("logid["+LogId+"] ", msg)
	}
	mylog.Println("logid["+LogId+"] ", msg)
}
