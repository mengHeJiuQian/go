package mlog

import (
	"fmt"
	"time"
)

// Logger日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

func NewLog(logLevel string) ConsoleLogger {
	level, err := parseLogLevel(logLevel)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(level LogLevel) bool {
	return level >= c.Level
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)

}

func (c ConsoleLogger) log(logLevel LogLevel, format string, a ...interface{}) {
	if c.enable(logLevel) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n",
			now.Format("2008-01-02 03:04:05"), getLogString(logLevel), fileName, funcName, lineNo, msg)
	}
}
