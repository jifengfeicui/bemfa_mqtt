package model

import (
	"bafa/Server"
	"bafa/global"
	"bafa/util"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-ini/ini"
)

type WolTopic struct {
	TopicName string
	Parameter *ini.Section
}

func (w *WolTopic) MessageHandler(client mqtt.Client, msg mqtt.Message) {
	msgStr := string(msg.Payload())
	global.Logger.Info("收到消息: " + msgStr)
	if msgStr == "on" {
		macAddrStr := w.Parameter.Key("mac").String()
		broadcastAddrStr := w.Parameter.Key("broadcast").String()
		Server.Wol(macAddrStr, broadcastAddrStr)
	}
}

func (w *WolTopic) Connect_mqtt() {
	util.Connect_mqtt(w.TopicName, w.MessageHandler)
}
