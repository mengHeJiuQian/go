package main

import (
	"hello-go/bili/laonanhai/mlog"
	"time"
)

// 日志库小作业

func main() {
	//log := mlog.NewLog("DEBUG")
	//
	//for {
	//    log.Debug("这是debug日志")
	//    log.Info("这是info日志%s", "刘洋")
	//    time.Sleep(time.Second * 1)
	//}

	log := mlog.NewFileLog("DEBUG", "./", "mainLog", 10*1024)

	for {
		log.Fatal("这是fatal日志")
		log.Error("这是error日志")
		log.Debug("这是debug日志")
		log.Info("这是info日志%s", "刘洋")
		time.Sleep(time.Second * 1)
	}
}
