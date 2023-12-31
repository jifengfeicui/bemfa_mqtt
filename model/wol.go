package model

import (
	"bafa/Server"
	"bafa/global"
	"bafa/util"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-ini/ini"
)

type WolTopic struct {
	TopicName string
	Parameter *ini.Section
}

func (w WolTopic) MessageHandler(_ mqtt.Client, msg mqtt.Message) {
	msgStr := string(msg.Payload())
	global.Logger.Info("收到消息: " + msgStr)
	if msgStr == "on" {
		macAddrStr := w.Parameter.Key("mac").String()
		broadcastAddrStr := w.Parameter.Key("broadcast").String()
		Server.Wol(macAddrStr, broadcastAddrStr)
	}
}

func (w WolTopic) ConnectMqtt() {
	util.ConnectMqtt(w.TopicName, w.MessageHandler)
}

func (w WolTopic) Verify() error {
	if !w.Parameter.HasKey("mac") {
		return fmt.Errorf("缺少mac")
	}
	if !w.Parameter.HasKey("broadcast") {
		return fmt.Errorf("缺少broadcast")
	}
	return nil
}
