package zaplog

import (
	"github.com/amsalt/glog"
	"github.com/amsalt/glog/adaptor"
	"go.uber.org/zap"
)

type logger struct {
	*adaptor.BaseAdaptor
	sugar  *zap.SugaredLogger
	logger *zap.Logger
}

func newLogger(opts *options, sugar *zap.SugaredLogger) *logger {
	l := new(logger)
	l.BaseAdaptor = adaptor.NewBaseAdaptor()
	if sugar != nil {
		l.sugar = sugar
	} else {
		panic("must assign a zap sugar logger.")
	}

	return l
}

func (l *logger) parseLevel(lv glog.Level) zap.AtomicLevel {
	l.BaseAdaptor.SetLevel(lv)
	switch lv {
	case glog.FatalLevel:
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	case glog.ErrorLevel:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case glog.WarningLevel:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case glog.InfoLevel:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	default:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	}
}

func (l *logger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.sugar.Debugf(format, args...)
}

func (l *logger) Warning(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *logger) Warningf(format string, args ...interface{}) {
	l.sugar.Warnf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.sugar.Infof(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}
func (l *logger) Errorf(format string, args ...interface{}) {
	l.sugar.Errorf(format, args...)
}

// Fatal output error at FatalLevel and panic
func (l *logger) Fatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

// Fatalf represents the printf format of Fatal.
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.sugar.Fatalf(format, args...)
}

type builder struct {
	logger *zap.SugaredLogger
}

func NewBuilder(logger *zap.SugaredLogger) glog.Builder {
	b := new(builder)
	b.logger = logger
	return b
}

func (b *builder) Build(opts ...glog.BuildOption) glog.Logger {
	defaultOpts := &options{}
	for _, o := range opts {
		o(defaultOpts)
	}
	return newLogger(defaultOpts, b.logger)
}
