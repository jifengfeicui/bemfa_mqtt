package model

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Topic interface {
	ConnectMqtt()
	MessageHandler(client mqtt.Client, msg mqtt.Message)
	Verify() error
}
