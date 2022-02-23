package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func Init() {
	writerSync := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "app-log.log",
		MaxSize:    500, // MB
		MaxBackups: 3,   // total file
		MaxAge:     28,  // hari
		LocalTime:  true,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writerSync,
		zap.InfoLevel,
	)
	logger = zap.New(core)
	// logger = log
}

func LogInfo(msg string) {
	logger.Info(msg)
}
