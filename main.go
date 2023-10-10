package main

import (
	"bafa/util"
)

func main() {
	RangeConfig()
	//阻塞主进程
	util.Block_main()
}
