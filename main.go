package main

import (
	"bafa/global"
	"bafa/messageHandler"
	"time"
)

func main() {
	// 读取INI文件
	topic := "goTest"
	go Connect_mqtt(topic, messageHandler.WolMessageHandler)
	for {
		global.Logger.Info("主进程正在运行...")
		time.Sleep(time.Minute)
	}
}
