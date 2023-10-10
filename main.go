package main

import (
	"bafa/global"
	"bafa/model"
	"fmt"
	"time"
)

func test() {

	topic := model.WolTopic{
		TopicName: "goTest",
		Parameter: global.Cfg.Section("goTest"),
	}
	topic.Connect_mqtt()
	block_main()
}

func block_main() {
	for {
		global.Logger.Info("主进程正在运行...")
		time.Sleep(time.Minute)
	}
}

func main() {
	// 遍历所有部分（sections）
	for _, section := range global.Cfg.Sections() {
		sectionName := section.Name()
		if sectionName == "DEFAULT" {
			continue
		}
		switch section.Key("struct").String() {
		case "wol":
			{
				topic := model.WolTopic{
					TopicName: sectionName,
					Parameter: section,
				}
				go topic.Connect_mqtt()
			}
		case "1":
			fmt.Println(1)
		case "2":
			fmt.Println(2)

		}

	}

	//阻塞主进程
	block_main()
}
