package logging

import (
	"flag"
	"fmt"
	"os"
	"service-notification/pkg/file"
)

var (
	LogSavePath = "/opt/log/service-notification/log/"
	LogSaveName = ""
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileName() string {
	flag.Parse()
//	return fmt.Sprintf("%s.%s", time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s.%s", "service-notification", LogFileExt)
}
func getDbLogFileName() string {
	return fmt.Sprintf("%s.%s", "db", LogFileExt)
}
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s", LogSaveName, getLogFileName())

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}


func openLogFile(fileName string, filePath string) (*os.File,error) {
	if perm := file.CheckPermission(filePath); perm == true {
		return nil, fmt.Errorf("file.CheckPermission denied src:%s", filePath)
	}
	if err := file.IsNotExistMkDir(filePath);err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src:%s, err:%v", filePath, err)
	}

	f, err := file.Open(filePath + fileName, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to open file file:%s%s,err:%v", filePath, fileName, err)
	}
	return f,nil
}

func mkDir(){
	err := os.MkdirAll(getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}