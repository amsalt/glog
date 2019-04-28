package log

import "sync"

// package log implements a simple logging facade for golang logger library.

// Level defines the log Level type as uint32.
type Level uint32

const (
	// FatalLevel only outuputs FatalLevel logs and panic.
	FatalLevel Level = 0

	// ErrorLevel only outuputs ErrorLevel and FatalLevel logs.
	ErrorLevel Level = 1

	// WarningLevel will ignore DebugLevel, InfoLevel log and output other levels of logs.
	WarningLevel Level = 2

	// InfoLevel only outuputs InfoLevel, WarningLevel, ErrorLevel and FatalLevel logs.
	InfoLevel Level = 3

	// DebugLevel will output all levels of logs.
	DebugLevel Level = 4
)

// ParseLevel parses the log level from string.
func ParseLevel(level string) Level {
	switch level {
	case "info":
		return InfoLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel

	default:
		return DebugLevel
	}
}

// Logger defines the logging facade interface.
type Logger interface {
	SetLevel(lv Level)
	GetLevel() Level

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	// Fatal output error at FatalLevel and panic
	Fatal(args ...interface{})
	// Fatalf represents the printf format of Fatal.
	Fatalf(format string, args ...interface{})
}

type BuildOption func(interface{})
type Builder interface {
	Build(opts ...BuildOption) Logger
}

var configuredBuilder Builder
var loggerInstance Logger
var once sync.Once

func init() {
	var defaultBuidler Builder
	SetBuilder(defaultBuidler)
}

func SetBuilder(builder Builder) {
	configuredBuilder = builder
}

func SetLogger(logger Logger) {
	loggerInstance = logger
}

func GetLogger() Logger {
	if loggerInstance == nil {
		once.Do(func() {
			loggerInstance = configuredBuilder.Build()
		})
	}
	return loggerInstance
}

func SetLevel(lv Level) {
	GetLogger().SetLevel(lv)
}
func GetLevel() Level {
	return GetLogger().GetLevel()
}

func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

func Warning(args ...interface{}) {
	GetLogger().Warning(args...)
}

func Warningf(format string, args ...interface{}) {
	GetLogger().Warningf(format, args...)
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}
func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}
func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

// Fatal output error at FatalLevel and panic
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// Fatalf represents the printf format of Fatal.
func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}
