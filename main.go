package main

import (
	"bafa/initialize"
	"bafa/util"
)

func main() {
	initialize.InitLogger()
	RangeConfig()
	//阻塞主进程
	util.Block_main()
}
