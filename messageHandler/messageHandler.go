package messageHandler

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MessageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("收到消息：%s\n", msg.Payload())
	// 在这里处理收到的消息，可以根据需要进行自定义处理
}
