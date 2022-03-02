package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File
	DefaultPrefix  = ""
	DefaultCallDepth = 2

	logger *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

func setPrefix(level Level) {
	pc, file, line, ok := runtime.Caller(DefaultCallDepth)
	funcName := runtime.FuncForPC(pc).Name()
	fileName := filepath.Base(file)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%s:%d]", levelFlags[level],fileName, funcName, line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}