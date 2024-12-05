package initialize

import (
	"bafa/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	Logger *zap.Logger
	//SugarLogger *zap.SugaredLogger
)

func InitLogger() {
	writeSyncer := getLogWriter()
	consoleWriteSyncer := zapcore.AddSync(os.Stdout) // 添加命令行输出配置
	encoder := getEncoder()
	// 创建多输出的 WriteSyncer
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(writeSyncer, consoleWriteSyncer)
	core := zapcore.NewCore(encoder, multiWriteSyncer, zapcore.DebugLevel)
	//core := zapcore.NewCore(encoder, multiWriteSyncer, zapcore.InfoLevel)
	Logger = zap.New(core, zap.AddCaller())
	global.SugarLogger = Logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./stdout.log")
	return zapcore.AddSync(file)
}
