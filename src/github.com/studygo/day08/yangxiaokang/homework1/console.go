package homework1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ConsoleLogger 定义结构体
type ConsoleLogger struct {
	level    logLevel
	isFile   bool
	filePath string
	fileName string
	fileObj  *os.File
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(level string) *ConsoleLogger {
	level = strings.ToUpper(level)
	Level, err := parseLogLevel(level)
	if err != nil {
		fmt.Printf("解析级别出现错误，err:%v\n", err)
	}
	return &ConsoleLogger{
		level:  Level,
		isFile: false,
	}
}

//日志开关，确定当前日志级别是否可以打印，返回true则可以打印
func (c *ConsoleLogger) isTrue(cLevel logLevel) bool {
	return cLevel >= c.level
}

//定义日志格式，打印日志到文件或屏幕
func (c *ConsoleLogger) messageFormat(lv logLevel, message string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	//定义日志格式
	msg := fmt.Sprintf("[%s] [%v] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, message)
	if c.isTrue(lv) {
		fmt.Printf(msg) //输出至屏幕
		if c.isFile {
			if c.checkTime() {
				c.MoveFile()
			}
		}
		_, err := fmt.Fprintf(c.fileObj, msg)
		if err != nil {
			fmt.Println("err:", err)
		}
	}
}

// Debug ...
func (c *ConsoleLogger) Debug(message string) {
	c.messageFormat(DEBUG, message)
}

// Info ...
func (c *ConsoleLogger) Info(message string) {
	c.messageFormat(INFO, message)
}
