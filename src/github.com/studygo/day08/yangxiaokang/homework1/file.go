package homework1

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileInit 如用户选择输出至文件，初始化文件
func (c *ConsoleLogger) FileInit(isFile bool, filePath, fileName string) error {
	c.isFile = isFile
	c.filePath = filePath
	c.fileName = fileName
	//fmt.Println(c.isFile,isFile,filePath) //为毛不打印
	logPath := path.Join(filePath, fileName)
	fileObj, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open logpath failed,err:%v \n", err)
		return err
	}
	c.fileObj = fileObj
	return nil
}

// checkTime 检查当前时间是否为每分钟的00秒
func (c *ConsoleLogger) checkTime() bool {
	// timef := time.Now().Format("200601021504") + "00"
	// checkTime, err := strconv.Atoi(timef)
	// nowTime, err := strconv.Atoi(time.Now().Format("20060102150405"))
	// if err != nil {
	// 	fmt.Printf("转换错误：%s", err)
	// }
	// //fmt.Printf("时间 %v %T,标志时间:%v %T \n",nowTime,nowTime,checkTime,checkTime)
	// return checkTime == nowTime

	return time.Now().Format("05") == "00"

}

// MoveFile 切割日志文件
func (c *ConsoleLogger) MoveFile() {
	c.fileObj.Close()
	nowStr := time.Now().Format("20060102150405")
	logPath := path.Join(c.filePath, c.fileName)
	newLogPath := logPath + nowStr
	os.Rename(logPath, newLogPath)

	fileObj, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开新的日志文件错误,%s", err)
		return
	}
	c.fileObj = fileObj // 没有生效?
}
