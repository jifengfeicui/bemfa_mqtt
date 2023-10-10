package main

import (
	"bafa/global"
	"bafa/model"
	"bafa/util"
	"fmt"
)

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
				err := topic.Verify()
				if err != nil {
					global.Logger.Error(err.Error())
					break
				}
				go topic.Connect_mqtt()
			}
		case "test":
			fmt.Println("test")
		}
	}

	//阻塞主进程
	util.Block_main()
}
