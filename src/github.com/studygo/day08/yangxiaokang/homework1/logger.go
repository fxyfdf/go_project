package homework1

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"errors"
)

type logLevel uint16

//定义日志级别
const (
	UNKNOW logLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//获取信息
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

//解析用户输入的日志级别
func parseLogLevel(level string) (logLevel, error) {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效输入")
		return UNKNOW, err
	}

}

func getLogString(lv logLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}
