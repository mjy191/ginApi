package myLogger

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

// LogWrite 日志，实现了Write接口
var LogWrite *logWrite

type logWrite struct {
	t  time.Time
	fp *os.File
	m  sync.RWMutex
}

// 返回文件写日志文件指针
func (l *logWrite) GetFp() *os.File {
	return l.fp
}

// init 创建runtime目录，并初始化Logger
func init() {
	if !isDir("logs") {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			panic("无法创建logs目录")
		}
	}

	LogWrite = new()
}

// Write 实现Write接口，用于写入
func (l *logWrite) Write(p []byte) (n int, err error) {
	today := dateToStr(time.Now())
	loggerDate := dateToStr(l.t)

	//如果当前日期与logger日期不一致，表示是新的一天，需要关闭原日志文件，并更新日期与日志文件
	if today != loggerDate && l.fp != nil {
		l.fp.Close()
		l.fp = nil
	}

	if l.fp == nil {
		l.setLogfile()
	}

	//写入
	if l.fp != nil {
		return l.fp.Write(p)
	}

	return 0, errors.New("无法写入日志")
}

// new 初始化
func new() *logWrite {
	l := &logWrite{
		t: time.Now(),
	}

	l.setLogfile()
	return l
}

// setLogfile 更新日志文件
func (l *logWrite) setLogfile() error {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	hour := timeNow.Hour()
	dir := fmt.Sprintf("logs/%d%02d", year, month)

	//锁住，防止并发时，多次执行创建。os.MkdirAll在目录存在时，也不会返回错误，锁不锁都行
	l.m.Lock()
	defer l.m.Unlock()
	if !isDir(dir) {
		err := os.MkdirAll(dir, 0755)
		return err
	}

	logfile := fmt.Sprintf("%s/%d%02d%02d%02d.log", dir, year, month, day, hour)
	//打开新的日志文件，用于写入
	fp, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	l.fp = fp
	return nil
}

// isDir 是否是目录
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// dateToStr 时间转换为日期字符串
func dateToStr(t time.Time) string {
	return t.Format("2006-01-02 15")
}
