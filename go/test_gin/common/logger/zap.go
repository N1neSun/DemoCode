package logger

import (
	"test_gin/common/setting"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger Log *zap.SugaredLogger
type Logger struct {
	Log *zap.SugaredLogger
}

type ILogger interface {
	////init 初始化日志配置
	Init()
	//Info 打印信息
	Info(args ...interface{})
	//Infof 打印信息，附带template信息
	Infof(template string, args ...interface{})
	//Warn 打印警告
	Warn(args ...interface{})
	//Warnf 打印警告，附带template信息
	Warnf(template string, args ...interface{})
	//Error 打印错误
	Error(args ...interface{})
	//Errorf 打印错误，附带template信息
	Errorf(template string, args ...interface{})
	//Panic 打印Panic信息
	Panic(args ...interface{})
	//Panicf 打印Panic信息，附带template信息
	Panicf(template string, args ...interface{})
	//DPanic 打印DPanic信息，附带template信息
	DPanic(args ...interface{})
	//DPanicf 打印DPanic信息
	DPanicf(template string, args ...interface{})
}

func NewLog() ILogger {
	return &Logger{}
}

//init 初始化日志配置
func (l *Logger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  setting.Config.APP.LogPath,
		MaxSize:   1024, //MB
		LocalTime: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.NewAtomicLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	l.Log = logger
}

//Info 打印信息
func (l *Logger) Info(args ...interface{}) {
	l.Log.Info(args...)
}

//Infof 打印信息，附带template信息
func (l *Logger) Infof(template string, args ...interface{}) {
	l.Log.Infof(template, args...)
}

//Warn 打印警告
func (l *Logger) Warn(args ...interface{}) {
	l.Log.Warn(args...)
}

//Warnf 打印警告，附带template信息
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.Log.Warnf(template, args...)
}

//Error 打印错误
func (l *Logger) Error(args ...interface{}) {
	l.Log.Error(args...)
}

//Errorf 打印错误，附带template信息
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.Log.Errorf(template, args...)
}

//Panic 打印Panic信息
func (l *Logger) Panic(args ...interface{}) {
	l.Log.Panic(args...)
}

//Panicf 打印Panic信息，附带template信息
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.Log.Panicf(template, args...)
}

//DPanic 打印Panic信息
func (l *Logger) DPanic(args ...interface{}) {
	l.Log.DPanic(args...)
}

//DPanicf 打印DPanic信息，附带template信息
func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.Log.DPanicf(template, args...)
}
