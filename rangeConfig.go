package main

import (
	"bafa/global"
	"bafa/model"
	"fmt"
	"github.com/go-ini/ini"
)

func RangeConfig() {
	// 遍历配置文件所有部分（sections）
	for _, section := range global.Cfg.Sections() {
		sectionName := section.Name()
		if sectionName == "DEFAULT" {
			continue
		}
		//注册服务
		registerServer(section)
	}
}

func registerServer(section *ini.Section) {
	switch section.Key("struct").String() {
	case "wol":
		{
			topic := model.WolTopic{
				TopicName: section.Name(),
				Parameter: section,
			}
			if run(topic) != nil {
				break
			}
		}
	case "test":
		fmt.Println(section.Name(), "test")
	default:
		fmt.Println(section.Name(), section.Key("struct").String())
	}

}

func run(topic model.Topic) error {
	err := topic.Verify()
	if err != nil {
		global.SugarLogger.Error(err.Error())
		return err
	}
	go topic.ConnectMqtt()
	return nil
}
