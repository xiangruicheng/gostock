package server

import (
	"log"
)

const (
	LogLevelInfo  = "INFO"
	LogLevelDebug = "DEBUG"
	LogLevelError = "ERROR"
)

func Log(level string, msg string) {
	//设置输出项
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//设置前缀：日志级别
	log.SetPrefix("[" + level + "] ")
	//设置日志文件
	//logFile, err := os.OpenFile("./output/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	fmt.Println("open log file failed, err:", err)
	//	return
	//}
	//log.SetOutput(logFile)

	//打印日志
	log.Println(msg)
}
