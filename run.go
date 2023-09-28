package main

import (
	"lucky/common/logger"
	"lucky/spider"
	"time"
)

func main() {
	//初始化日志模块
	logger.InitLogger()

	//开启一个定时器，每五分钟执行一次
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-ticker.C:
			// TODO: 执行爬虫任务的代码
			spider.Spider()
		}
	}

}
