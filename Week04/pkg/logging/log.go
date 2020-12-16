package logging

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix = ""
	DefaultCallerDepth = 2

	Logger *log.Logger
	DbLogger gorm.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = openLogFile(fileName, filePath)
	if err != nil {
		panic("Setup openLogFile err:" + err.Error())
	}
	Logger = log.New(F, DefaultPrefix, log.LstdFlags)
	dbFileName := getDbLogFileName()
	FDb, err := openLogFile(dbFileName, filePath)
	if err != nil {
		panic("Setup openLogFile err:" + err.Error())
	}
	DbLogger = gorm.Logger{log.New(FDb, "\r\n", log.LstdFlags)}
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	Logger.Println(v)
}
func Info(v ...interface{}) {

	setPrefix(INFO)
	Logger.Println(v)
}
func Infof(format string, v ...interface{})  {
	setPrefix(INFO)
	Logger.Printf(format, v...)
}
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	Logger.Println(v)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	Logger.Println(v)
}
func Errorf(format string, v ...interface{})  {
	setPrefix(ERROR)
	Logger.Printf(format, v...)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	Logger.Println(v)
}
func setPrefix(level Level)  {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		if levelFlags[level] == "WARN" || levelFlags[level] == "ERROR" || levelFlags[level] == "FATAL" {
			logPrefix = fmt.Sprintf("\033[31m[%s]\033[0m\033[35m[%s:%d]\033[0m", levelFlags[level], filepath.Base(file), line)
		}else {
			logPrefix = fmt.Sprintf("\033[37m[%s]\033[0m\033[35m[%s:%d]\033[0m", levelFlags[level], filepath.Base(file), line)
		}
	}else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	Logger.SetPrefix(logPrefix)
}