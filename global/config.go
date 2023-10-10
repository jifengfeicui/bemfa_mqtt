package global

import (
	"github.com/go-ini/ini"
	"log"
)

var Cfg *ini.File
var BemfaBroker string
var BemfaPort string
var Bemfa_client_id string

func init() {
	var err error
	// 读取INI文件
	Cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件: %v", err)
	}

	// 获取配置项的值
	mainSection := Cfg.Section("DEFAULT") // 选择INI文件中的某个section
	BemfaBroker = mainSection.Key("bemfa_broker").String()
	BemfaPort = mainSection.Key("bemfa_port").String()
	Bemfa_client_id = mainSection.Key("bemfa_client_id").String()
	if err != nil {
		log.Fatalf("无法解析 bemfa_port: %v", err)
	}

	//// 遍历所有部分（sections）
	//for _, section := range Cfg.Sections() {
	//	sectionName := section.Name()
	//	if sectionName == "DEFAULT" {
	//		continue
	//	}
	//	fmt.Printf("topic名称：%s\n", sectionName)
	//
	//	// 遍历部分中的所有键值对
	//	for _, key := range section.Keys() {
	//		keyName := key.Name()
	//		keyValue := key.String()
	//		fmt.Printf("键：%s，值：%s\n", keyName, keyValue)
	//	}
	//
	//}
}
