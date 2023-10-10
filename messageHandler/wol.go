package messageHandler

import (
	"bafa/global"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"net"
)

func Wol(macAddrStr, broadcastAddrStr string) {

	// 解析MAC地址
	macAddr, err := net.ParseMAC(macAddrStr)
	if err != nil {
		global.Logger.Error("无效的MAC地址:", zap.Error(err))
	}

	// 解析广播地址
	broadcastAddr, err := net.ResolveUDPAddr("udp", broadcastAddrStr+":9")
	if err != nil {
		global.Logger.Error("无效的广播地址:", zap.Error(err))
	}

	// 创建Magic Packet
	magicPacket := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	for i := 0; i < 16; i++ {
		magicPacket = append(magicPacket, macAddr...)
	}

	// 创建UDP连接并发送Magic Packet
	conn, err := net.DialUDP("udp", nil, broadcastAddr)
	if err != nil {
		global.Logger.Error("无效的广播地址:", zap.Error(err))
	}
	defer conn.Close()

	_, err = conn.Write(magicPacket)
	if err != nil {
		global.Logger.Error("无法发送Magic Packet:", zap.Error(err))
	}
	global.Logger.Info("Magic Packet发送成功！")
}

func WolMessageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("收到消息：%s\n", msg.Payload())
	if string(msg.Payload()) == "on" {
		section := global.Cfg.Section("goTest")
		macAddrStr := section.Key("mac").String()
		broadcastAddrStr := section.Key("broadcast").String()
		Wol(macAddrStr, broadcastAddrStr)
	}
}
