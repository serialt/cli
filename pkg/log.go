package pkg

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugar *zap.SugaredLogger

type Logger struct {
	LogLevel      string // 日志级别
	LogFile       string // 日志文件存放路径,如果为空，则输出到控制台
	LogType       string // 日志类型，支持 txt 和 json ，默认txt
	LogMaxSize    int    //单位M
	LogMaxBackups int    // 日志文件保留个数
	LogMaxAge     int    // 单位天
	LogCompress   bool   // 压缩轮转的日志
	LogColor      bool   // 日志级别分颜色
}

func LevelToZapLevel(level string) zapcore.Level {
	// 转换日志级别
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}

}

func (lg *Logger) NewMyLogger() *zap.Logger {

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(LevelToZapLevel(lg.LogLevel))

	// 输出的消息
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 设置大写日志级别无颜色
		EncodeTime:     zapcore.RFC3339TimeEncoder,     // 日期格式 2022-07-23T10:49:47+08:00
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行消耗的时间转化成浮点型的秒
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 以包/文件:行号 格式化调用堆栈
	}
	// 简易设置日志级别分颜色
	if lg.LogColor {
		// capital, capitalColor, color, default LowercaseLevelEncoder
		encoderConfig.EncodeLevel.UnmarshalText([]byte("capitalColor"))
	}

	// 日志输出类型
	var encoder zapcore.Encoder
	switch lg.LogType {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)

	}

	var core zapcore.Core
	if lg.LogFile == "" {
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), atomicLevel)
	} else {
		file := zapcore.AddSync(&lumberjack.Logger{
			Filename:   lg.LogFile,       // 日志文件
			MaxSize:    lg.LogMaxSize,    // 单个日志文件大小，单位M
			MaxBackups: lg.LogMaxBackups, // 轮转保留个数
			MaxAge:     lg.LogMaxAge,     // 最长保留时间，单位天
			Compress:   lg.LogCompress,   // 默认不压缩
		})
		core = zapcore.NewCore(encoder, zapcore.AddSync(file), atomicLevel)
	}

	// 开启开发模式，堆栈跟踪: [zap.AddCaller()]
	myLogger := zap.New(core, zap.AddCaller())
	return myLogger
}

// NewLogger 自定日志配置可以参考此方法
func NewLogger(logLevel, logFile, logtype string, logColor bool) *zap.Logger {
	lg := &Logger{
		LogLevel:      logLevel,
		LogFile:       logFile,
		LogType:       logtype,
		LogMaxSize:    50,
		LogMaxBackups: 10,
		LogMaxAge:     365,
		LogCompress:   true,
		LogColor:      logColor,
	}
	return lg.NewMyLogger()
}

// NewSugarLogger 创建一个sugar
func NewSugarLogger(logLevel, logFile, logType string, logColor bool) *zap.SugaredLogger {
	sugarLog := NewLogger(logLevel, logFile, logType, logColor)
	return sugarLog.Sugar()

}
