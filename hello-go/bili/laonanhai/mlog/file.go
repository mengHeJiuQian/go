package mlog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 记录日志文件相关功能

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	// 日志文件都打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func NewFileLog(logLevel, fp, fn string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(logLevel)
	if err != nil {
		panic(err)
	}
	fileLog := &FileLogger{
		Level:    level,
		filePath: fp,
		fileName: fn,

		maxFileSize: maxSize,
	}
	err = fileLog.initFile()
	if err != nil {
		panic(err)
	}
	return fileLog
}

func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	if f.enable(DEBUG) {
		f.log(DEBUG, format, a...)
	}
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		f.log(INFO, format, a...)
	}
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		f.log(ERROR, format, a...)
	}
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	if f.enable(FATAL) {
		f.log(FATAL, format, a...)
	}
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err=%v\n", err)
		return false
	}
	// 如果当前日志文件大小大于日志文件的最大值，返回TRUE
	return fileInfo.Size() > f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 切割日志文件步骤
	// 1.关闭文件流
	f.fileObj.Close()
	// 2.日志文件重命名为.bak后缀
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err=%v\n", err)
		return nil, err
	}
	nowStr := time.Now().Format("20060102150405000")
	logName := fileInfo.Name()
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	os.Rename(logName, newLogName)
	// 3.创建新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new file failed, err=%v\n", err)
		return nil, err
	}
	// 4.将新的文件对象赋给fileObj
	//f.fileObj = fileObj
	return fileObj, nil
}

func (f *FileLogger) log(logLevel LogLevel, format string, a ...interface{}) {
	if f.enable(logLevel) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n",
			now.Format("2008-01-02 03:04:05"), getLogString(logLevel), fileName, funcName, lineNo, msg)

		if logLevel >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n",
				now.Format("2008-01-02 03:04:05"), getLogString(logLevel), fileName, funcName, lineNo, msg)
		}
	}
}
