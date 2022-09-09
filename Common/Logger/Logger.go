package Logger

import (
	"fmt"
	"ginApi/Common/Tools"
	"log"
)

var Logid string
var mylog *log.Logger
var env bool = false

func init() {
	mylog = log.New(LogWrite, "", log.LstdFlags)
	if Tools.Config.Get("env") == "product" {
		env = true
	}
}

func Printf(msg string, v ...any) {
	if v != nil {
		msg = fmt.Sprintf(msg, v)
	}
	if !env {
		fmt.Println("logid["+Logid+"] ", msg)
	}
	mylog.Println("logid["+Logid+"] ", msg)
}

func Println(msg string) {
	if !env {
		fmt.Println("logid["+Logid+"] ", msg)
	}
	mylog.Println("logid["+Logid+"] ", msg)
}
