package Logger

import (
	"fmt"
	"log"
)

var Logid string
var mylog *log.Logger

func init() {
	mylog = log.New(LogWrite, "", log.LstdFlags)
}

func Printf(msg string, v ...any) {
	if v != nil {
		msg = fmt.Sprintf(msg, v)
	}
	fmt.Println("logid["+Logid+"] ", msg)
	mylog.Println("logid["+Logid+"] ", msg)
}

func Println(msg string) {
	fmt.Println("logid["+Logid+"] ", msg)
	mylog.Println("logid["+Logid+"] ", msg)
}
