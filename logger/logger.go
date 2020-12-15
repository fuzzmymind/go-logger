package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config struct to set up logging by the application
type Config struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LogLevel   zapcore.Level
}

const (
	// DebugLevel log level
	DebugLevel = zap.InfoLevel
	// InfoLevel log level
	InfoLevel = zap.InfoLevel
	// WarnLevel log level
	WarnLevel = zap.WarnLevel
	// ErrorLevel log level
	ErrorLevel = zap.ErrorLevel
	// FatalLevel log level
	FatalLevel = zap.FatalLevel
)

var (
	log *zap.Logger
)

// Init function that sets up the logging parameters for the application
func Init(lc Config) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   lc.FileName,
		MaxSize:    lc.MaxSize,
		MaxBackups: lc.MaxBackups,
		MaxAge:     lc.MaxAge,
	})
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "time",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		lc.LogLevel,
	)

	log = zap.New(core)
}

// Field function for any custom key value pairs that needs to be logged
func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// Info function to log information logs
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

// Error function to log error logs
func Error(msg string, tags ...zap.Field) {
	log.Error(msg, tags...)
	log.Sync()
}

// Debug function to log debug logs
func Debug(msg string, tags ...zap.Field) {
	log.Debug(msg, tags...)
	log.Sync()
}

// Warn function to log warning logs
func Warn(msg string, tags ...zap.Field) {
	log.Warn(msg, tags...)
	log.Sync()
}

// Fatal function to log fatal logs. This will exit the program.
func Fatal(msg string, tags ...zap.Field) {
	log.Fatal(msg, tags...)
	log.Sync()
}
