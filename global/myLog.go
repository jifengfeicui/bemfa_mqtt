package global

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer Logger.Sync() // 在程序退出时确保刷新所有缓冲的日志

}
