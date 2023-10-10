package main

import (
	"bafa/global"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Connect_mqtt(topic string, messageHandler func(client mqtt.Client, msg mqtt.Message)) {
	// 连接参数
	//brokerURL := "tcp://bemfa.com:9501"
	brokerURL := fmt.Sprintf("tcp://%s:%s", global.BemfaBroker, global.BemfaPort)
	clientID := global.Bemfa_client_id

	qos := 1 // 可根据需求调整

	// 创建 MQTT 客户端
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(clientID)
	opts.SetCleanSession(false)
	opts.SetDefaultPublishHandler(messageHandler) // 设置消息处理函数

	client := mqtt.NewClient(opts)

	// 连接到 MQTT 服务器
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("连接到 MQTT 服务器失败: %v", token.Error())
		return
	}

	// 订阅主题
	if token := client.Subscribe(topic, byte(qos), nil); token.Wait() && token.Error() != nil {
		log.Fatalf("订阅主题失败: %v", token.Error())
		return
	}

	fmt.Printf("已连接到 MQTT 服务器，订阅主题：%s\n", topic)

	// 持续监听消息
	for {
		time.Sleep(time.Second)
	}
}

//func messageHandler(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("收到消息：%s\n", msg.Payload())
//	// 在这里处理收到的消息，可以根据需要进行自定义处理
//}
