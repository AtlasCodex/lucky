package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger 初始化日志模块
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Log = zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	dir := filepath.Dir("log/app.log")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil
		}
	}
	file, _, _ := zap.Open("log/app.log")
	return zapcore.AddSync(file)
}

func init() {
	// 初始化日志模块
	InitLogger()
}
