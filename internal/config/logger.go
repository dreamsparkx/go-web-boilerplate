package config

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var AppLogger *zap.SugaredLogger

func createDirectoryIfNotExists() {
	path, _ := os.Getwd()
	if _, err := os.Stat(fmt.Sprintf("%s/logs", path)); os.IsNotExist(err) {
		_ = os.Mkdir("logs", os.ModePerm)
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/logs.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func InitLogger(inProduction bool) {
	if inProduction {
		createDirectoryIfNotExists()
		writerSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
		logger := zap.New(core, zap.AddCaller())
		AppLogger = logger.Sugar()
	} else {
		logger, _ := zap.NewProduction()
		AppLogger = logger.Sugar()
	}
}
