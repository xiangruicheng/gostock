package server

import (
	"gostock/config"
	"log"
	"os"
	"time"
)

type LogServer struct {
	logger *log.Logger
}

var Log *LogServer

func init() {
	Log = new(LogServer)
	Log.logger = new(log.Logger)
	Log.logger.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// create log dir
	_, err := os.Stat(config.Data.App.LogPath)
	if err != nil {
		err = os.Mkdir(config.Data.App.LogPath, 0755)
	}
}

func (l *LogServer) getFile() *os.File {
	filename := config.Data.App.LogPath + time.Now().Format("2006-01-02") + ".log"
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("open log file fail:" + filename)
	}
	l.logger.SetOutput(logFile)
	return logFile
}

func (l *LogServer) Info(msg string) {
	file := l.getFile()
	defer file.Close()
	l.logger.SetPrefix("[INFO] ")
	l.logger.Println(msg)
}

func (l *LogServer) Debug(msg string) {
	file := l.getFile()
	defer file.Close()
	l.logger.SetPrefix("[DEBUG] ")
	l.logger.Println(msg)
}

func (l *LogServer) Error(msg string) {
	file := l.getFile()
	defer file.Close()
	l.logger.SetPrefix("[ERROR] ")
	l.logger.Println(msg)
}
