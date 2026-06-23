package model

import (
	"bafa/global"
	"bafa/util"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-ini/ini"
	"os"
	"strings"
)

type ScreenRotateTopic struct {
	TopicName string
	Parameter *ini.Section
}

func (s ScreenRotateTopic) MessageHandler(_ mqtt.Client, msg mqtt.Message) {
	msgStr := string(msg.Payload())
	global.SugarLogger.Info("收到消息: " + msgStr)

	rotate, ok := rotateForPayload(msgStr)
	if !ok {
		global.SugarLogger.Info("忽略屏幕旋转消息: " + msgStr)
		return
	}

	if err := s.rotate(rotate); err != nil {
		global.SugarLogger.Error("屏幕旋转失败: " + err.Error())
	}
}

func (s ScreenRotateTopic) ConnectMqtt() {
	util.ConnectMqtt(s.TopicName, s.MessageHandler)
}

func (s ScreenRotateTopic) Verify() error {
	for _, key := range []string{"ssh_host", "ssh_user", "ssh_password"} {
		if !s.Parameter.HasKey(key) {
			return fmt.Errorf("缺少%s", key)
		}
	}
	script := s.Parameter.Key("script").MustString("tmp/rotatescreen.ps1")
	if _, err := os.Stat(script); err != nil {
		return fmt.Errorf("屏幕旋转脚本不可用: %w", err)
	}
	return nil
}

func (s ScreenRotateTopic) rotate(rotate int) error {
	scriptPath := s.Parameter.Key("script").MustString("tmp/rotatescreen.ps1")
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("读取屏幕旋转脚本失败: %w", err)
	}

	display := s.Parameter.Key("display").MustInt(2)
	command := fmt.Sprintf(`powershell -NoProfile -ExecutionPolicy Bypass -Command "$script = Join-Path $env:TEMP 'bemfa_rotatescreen.ps1'; [Console]::In.ReadToEnd() | Set-Content -LiteralPath $script -Encoding UTF8; & $script -Display %d -Rotate %d"`, display, rotate)
	// ponytail: re-send the tiny script on each trigger; cache remotely if latency matters.
	return util.SSHCommandWithStdin(
		s.Parameter.Key("ssh_host").String(),
		s.Parameter.Key("ssh_port").MustString("22"),
		s.Parameter.Key("ssh_user").String(),
		s.Parameter.Key("ssh_password").String(),
		command,
		string(script),
	)
}

func rotateForPayload(payload string) (int, bool) {
	switch strings.TrimSpace(payload) {
	case "on":
		return 90, true
	case "off":
		return 0, true
	default:
		return 0, false
	}
}
