package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s",prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch  {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatal("permission:%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to openFile :%v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	log.Println(dir)
	err := os.Mkdir(dir+"/"+getLogFilePath(),os.ModePerm)
	if err != nil {
		panic(err)
	}
}