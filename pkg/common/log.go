package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

func init() {
	var coreArr []zapcore.Core
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	colorEncoder := getColorEncoder()
	//日志级别
	//error级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	allPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev >= zap.DebugLevel
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer), highPriority)
	errorCore := zapcore.NewCore(colorEncoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), allPriority)
	coreArr = append(coreArr, errorFileCore)
	coreArr = append(coreArr, errorCore)
	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
}

func getColorEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/error.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
