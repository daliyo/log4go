package log4go

import (
	"log"
	"os"
)

const (
	//日志文件默认10MB
	logSize = 10
)

// fileLogger 文件日志记录器
type fileLogger struct{}

func (l *fileLogger) Write(p []byte) (n int, err error) {
	f, _ := os.OpenFile("app.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	f.Write(p)
	return len(p), nil
}

// InitFileLogger 初始化文件日志记录器
func InitFileLogger(path string) {
	log.SetOutput(new(fileLogger))
}
