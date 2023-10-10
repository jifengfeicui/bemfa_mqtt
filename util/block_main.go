package util

import (
	"bafa/global"
	"time"
)

func Block_main() {
	for {
		currentTime := time.Now()
		formattedTime := currentTime.Format("2006-01-02 15:04:05")
		global.Logger.Info("主进程正在运行...")
		global.Logger.Info("当前时间" + formattedTime)
		time.Sleep(time.Hour)
	}
}
