package main

import (
	"bafa/messageHandler"
	"fmt"
	"time"
)

func main() {
	// 读取INI文件
	topic := "goTest"
	go Connect_mqtt(topic, messageHandler.MessageHandler)
	for {
		fmt.Println("主进程正在运行...")
		time.Sleep(time.Minute)
	}
}
