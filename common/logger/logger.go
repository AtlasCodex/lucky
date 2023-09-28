package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger 初始化日志模块
func InitLogger() {
	loadConfig() // 从配置文件加载配置
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	Log = zap.New(core, zap.AddCaller())
}

func loadConfig() {
	viper.SetConfigName("config") // 配置文件名 (without extension)
	viper.SetConfigType("toml")   // 配置文件类型
	viper.AddConfigPath(".")
	// 配置文件路径（当前目录）
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		os.Exit(1)
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	logFile := viper.GetString("logger.log_file")
	dir := filepath.Dir(logFile)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			Log.Error("Error creating log directory", zap.Error(err))
			os.Exit(1)
		}
	}

	file, err := os.Create(logFile)
	if err != nil {
		Log.Error("Error creating log file", zap.Error(err))
		os.Exit(1)
	}

	return zapcore.AddSync(file)
}

func init() {
	// 初始化日志模块
	InitLogger()
}
