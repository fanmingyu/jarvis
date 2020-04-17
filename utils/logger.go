package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

var _, TimeZone = time.Now().Zone()

type LogWriter struct {
	Path	string
	Prefix	string
	File	*os.File
	CurrentDate	int64
}

var logFileWriter = LogWriter{}

func LoggerInit(logPath string) {
	logFileWriter.Path = logPath
	logFileWriter.Prefix = "normal_"
	log.SetOutput(&logFileWriter)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

//实现write接口
func (fileWriter *LogWriter) Write(p []byte) (int, error) {
	fileWriter.CheckDate()
	return fileWriter.File.Write(p)
}

//检测日期，隔天就要新建log文件
func (fileWriter *LogWriter) CheckDate() {
	var err error
	now := time.Now()
	data := (now.Unix() + int64(TimeZone)) / 86400

	if data == fileWriter.CurrentDate {
		return
	}

	fileWriter.File.Close()

	fileName := fmt.Sprintf("%v%v%v", fileWriter.Prefix, now.Format("20060102"), ".log")
	fileWriter.File, err = os.OpenFile(fileWriter.Path+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(fmt.Sprintf("open log file failed. err:%s, file:%s\n", err, fileName))
	}

	//软链
	symlinkName := fileWriter.Path + "/" + "normal.log"
	fileWriter.createSymlink(fileName, symlinkName)

	fileWriter.CurrentDate = data
	fmt.Printf("logger file create. file:%v, time:%v\n", fileName, time.Now())
}

//创建软链
func (l *LogWriter) createSymlink(fileName string, symlinkName string) {
	_, err := os.Lstat(symlinkName)
	if err == nil {
		os.Remove(symlinkName)
	}

	err = os.Symlink(fileName, symlinkName)
	if err != nil {
		fmt.Printf("symlink create failed. err:%s\n", err)
	}
}
