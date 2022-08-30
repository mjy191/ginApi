package Logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var Logid string

func init() {
	logFile := LogWrite
	multiWrite := io.MultiWriter(os.Stdout, logFile.fp)
	log.New(logFile, "", log.LstdFlags)
	log.SetOutput(multiWrite)
}

func Printf(msg string, v ...any) {
	if v != nil {
		msg = fmt.Sprintf(msg, v)
	}
	log.Println("logid["+Logid+"] ", msg)
}

func Println(msg ...any) {
	log.Println("logid["+Logid+"] ", msg)
}
