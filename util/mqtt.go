package util

import (
	"bafa/global"
	"fmt"
	"go.uber.org/zap"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConnectMqtt(topic string, messageHandler func(client mqtt.Client, msg mqtt.Message)) {
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
	opts.SetDefaultPublishHandler(messageHandler)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		global.SugarLogger.Error("MQTT 连接断开: ", zap.Error(err))
	})

	client := mqtt.NewClient(opts)

	// 连接到 MQTT 服务器
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		global.SugarLogger.Error("连接到 MQTT 服务器失败", zap.Error(token.Error()))
		return
	}

	// 订阅主题
	if token := client.Subscribe(topic, byte(qos), nil); token.Wait() && token.Error() != nil {
		global.SugarLogger.Error("订阅主题失败", zap.Error(token.Error()))
		return
	}

	global.SugarLogger.Info("已连接到 MQTT 服务器,订阅主题: " + topic)
	<-make(chan int)
}
