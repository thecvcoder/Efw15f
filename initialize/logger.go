package initialize

import (
	"fmt"
	"github.com/thecvcoder/cloud-api-go/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

/*
*	zap提供两种日志记录器，分别是 Logger和SugaredLogger
*	Logger性能更好，更快，但只支持强类型的结构化日志
*	SugaredLogger支持结构化和printf风格的日志记录，用法也几乎相同。
*	详见：https://pkg.go.dev/go.uber.org/zap#hdr-Choosing_a_Logger
 */

var log *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	now := time.Now()
	// info日志相关
	infoLogFileName := fmt.Sprintf("%s/info/%04d-%02d-%02d.log", global.CONFIG.Logs.Path, now.Year(), now.Month(), now.Day())
	// 错误日志
	errorLogFileName := fmt.Sprintf("%s/error/%04d-%02d-%02d.log", global.CONFIG.Logs.Path, now.Year(), now.Month(), now.Day())
	var coreArr []zapcore.Core

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "file",
		FunctionKey:   "func",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zap.ErrorLevel && level >= zap.DebugLevel
	})

	// 当yml配置中的等级大于Error时，lowPriority级别日志停止记录
	if global.CONFIG.Logs.Level >= 2 {
		lowPriority = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return false
		})
	}

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   infoLogFileName,               //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    global.CONFIG.Logs.MaxSize,    //文件大小限制,单位MB
		MaxAge:     global.CONFIG.Logs.MaxAge,     //日志文件保留天数
		MaxBackups: global.CONFIG.Logs.MaxBackups, //最大保留日志文件数量
		LocalTime:  false,
		Compress:   global.CONFIG.Logs.Compress, //是否压缩处理
	})
	// 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)

	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   errorLogFileName,              //日志文件存放目录
		MaxSize:    global.CONFIG.Logs.MaxSize,    //文件大小限制,单位MB
		MaxAge:     global.CONFIG.Logs.MaxAge,     //日志文件保留天数
		MaxBackups: global.CONFIG.Logs.MaxBackups, //最大保留日志文件数量
		LocalTime:  false,
		Compress:   global.CONFIG.Logs.Compress, //是否压缩处理
	})
	// 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)

	logger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
	log = logger.Sugar()
	log.Info("初始化zap日志完成!")

	return log
}
