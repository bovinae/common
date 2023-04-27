package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once sync.Once
	l    *zap.Logger
)

type options struct {
	fileName string // Default empty, if fileName is empty, print to console.
	level    zapcore.Level
}

type Option func(opt *options)

func WithFileName(fileName string) Option {
	return func(opt *options) {
		opt.fileName = fileName
	}
}

func WithLevel(level zapcore.Level) Option {
	return func(opt *options) {
		opt.level = level
	}
}

func InitLogger(opts ...Option) *zap.Logger {
	if l != nil {
		return l
	}

	opt := options{
		level: zap.InfoLevel,
	}
	for _, o := range opts {
		o(&opt)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	var hook *lumberjack.Logger
	if len(opt.fileName) > 0 {
		hook = &lumberjack.Logger{
			Filename:   opt.fileName, // 日志文件路径 "./logs/spikeProxy1.log"
			MaxSize:    128,          // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 30,           // 日志文件最多保存多少个备份
			MaxAge:     30,           // 文件最多保存多少天
			Compress:   true,         // 是否压缩
		}
	}
	var writeSyncer zapcore.WriteSyncer
	if hook == nil {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)) // 打印到控制台
	} else {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(hook)) // 打印到文件
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(opt.level)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		writeSyncer,
		atomicLevel, // 日志级别
	)
	// 设置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	once.Do(func() {
		// 构造日志
		l = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development())
	})
	return l
}

func Debug(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

func With(fields ...zap.Field) {
	l.With(fields...)
}
