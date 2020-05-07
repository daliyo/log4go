package log4go

import (
	"fmt"
	"log"
)

// Level 日志等级
type Level int

const (
	// LevelError 错误
	LevelError Level = 1 << iota
	// LevelWarn 警告
	LevelWarn
	// LevelInfo 消息
	LevelInfo
	// LevelDebug 调试
	LevelDebug
)

const (
	prefixError = "ERROR "
	prefixWarn  = "WARN  "
	prefixInfo  = "INFO  "
	prefixDebug = "DEBUG "
)

var defaultLevel = LevelInfo | LevelWarn | LevelError

// SetLevel 设置日志等级
func SetLevel(target Level) {
	defaultLevel = target
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func enabled(target Level) bool {
	return target&defaultLevel != 0
}

// Error 打印一个错误
func Error(v ...interface{}) {
	if enabled(LevelError) {
		log.Output(2, fmt.Sprint(prefixError, v))
	}
}

// Errorf 打印一个错误
func Errorf(format string, v ...interface{}) {
	if enabled(LevelError) {
		log.Output(2, prefixError+fmt.Sprintf(format, v...))
	}
}

// Warn 打印一个警告
func Warn(v interface{}) {
	if enabled(LevelWarn) {
		log.Output(2, fmt.Sprint(prefixWarn, v))
	}
}

// Warnf 打印一个警告
func Warnf(format string, v ...interface{}) {
	if enabled(LevelWarn) {
		log.Output(2, prefixWarn+fmt.Sprintf(format, v...))
	}
}

// Info 打印一条消息
func Info(v interface{}) {
	if enabled(LevelInfo) {
		log.Output(2, fmt.Sprint(prefixInfo, v))
	}
}

// Infof 打印一条消息
func Infof(format string, v ...interface{}) {
	if enabled(LevelInfo) {
		log.Output(2, prefixInfo+fmt.Sprintf(format, v...))
	}
}

// Debug 打印调试信息
func Debug(v ...interface{}) {
	if enabled(LevelDebug) {
		log.Output(2, fmt.Sprint(prefixDebug, v))
	}
}

// Debugf 打印调试信息
func Debugf(format string, v ...interface{}) {
	if enabled(LevelDebug) {
		log.Output(2, prefixDebug+fmt.Sprintf(format, v...))
	}
}
