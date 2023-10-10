package model

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Topic interface {
	Connect_mqtt()
	MessageHandler(client mqtt.Client, msg mqtt.Message)
}
