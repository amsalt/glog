package log

import (
	"path"
	"runtime"
	"strconv"
	"sync"
)

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

	SetPrintCallerLevel(lv Level)
	GetPrintCallerLevel() Level

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

func SetPrintCallerLevel(lv Level) {
	GetLogger().SetPrintCallerLevel(lv)
}

func GetPrintCallerLevel() Level {
	return GetLogger().GetPrintCallerLevel()
}

func Debug(args ...interface{}) {
	output(DebugLevel, args...)
}
func Debugf(format string, args ...interface{}) {
	outputf(DebugLevel, format, args...)
}

func Warning(args ...interface{}) {
	output(WarningLevel, args...)
}

func Warningf(format string, args ...interface{}) {
	outputf(WarningLevel, format, args...)
}

func Info(args ...interface{}) {
	output(InfoLevel, args...)
}
func Infof(format string, args ...interface{}) {
	outputf(InfoLevel, format, args...)
}

func Error(args ...interface{}) {
	output(ErrorLevel, args...)
}
func Errorf(format string, args ...interface{}) {
	outputf(ErrorLevel, format, args...)
}

// Fatal output error at FatalLevel and panic
func Fatal(args ...interface{}) {
	output(FatalLevel, args...)
}

// Fatalf represents the printf format of Fatal.
func Fatalf(format string, args ...interface{}) {
	outputf(FatalLevel, format, args...)
}

func output(level Level, args ...interface{}) {
	if GetLevel() < level {
		return
	}
	if GetPrintCallerLevel() <= level {
		args = append([]interface{}{addCallerInfo()}, args...)
	}
	switch level {
	case FatalLevel:
		GetLogger().Fatal(args...)
	case ErrorLevel:
		GetLogger().Error(args...)
	case WarningLevel:
		GetLogger().Warning(args...)
	case InfoLevel:
		GetLogger().Info(args...)
	case DebugLevel:
		GetLogger().Debug(args...)
	}
}

func outputf(level Level, format string, args ...interface{}) {
	if GetLevel() < level {
		return
	}
	if GetPrintCallerLevel() <= level {
		format = addCallerInfo() + format
	}
	switch level {
	case FatalLevel:
		GetLogger().Fatalf(format, args...)
	case ErrorLevel:
		GetLogger().Errorf(format, args...)
	case WarningLevel:
		GetLogger().Warningf(format, args...)
	case InfoLevel:
		GetLogger().Infof(format, args...)
	case DebugLevel:
		GetLogger().Debugf(format, args...)
	}
}

func addCallerInfo() (info string) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	info = "[" + filename + ":" + strconv.Itoa(line) + "] "
	return
}
