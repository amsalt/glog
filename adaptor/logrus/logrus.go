package logrus

import (
	"github.com/Sirupsen/logrus"
	"github.com/amsalt/log"
	"github.com/amsalt/log/adaptor"
)

type logger struct {
	*adaptor.BaseAdaptor
	entry *logrus.Entry
}

func (l *logger) SetLevel(lv log.Level) {
	l.BaseAdaptor.SetLevel(lv)
	switch lv {
	case log.FatalLevel:
		l.entry.Logger.SetLevel(logrus.FatalLevel)
	case log.ErrorLevel:
		l.entry.Logger.SetLevel(logrus.ErrorLevel)
	case log.WarningLevel:
		l.entry.Logger.SetLevel(logrus.WarnLevel)
	case log.InfoLevel:
		l.entry.Logger.SetLevel(logrus.InfoLevel)
	default:
		l.entry.Logger.SetLevel(logrus.DebugLevel)
	}
}

func (l *logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logger) Warning(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *logger) Warningf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}
func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Fatal output error at FatalLevel and panic
func (l *logger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

// Fatalf represents the printf format of Fatal.
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func newLogger(opts *options, lo *logrus.Logger) *logger {
	l := &logger{}
	l.BaseAdaptor = adaptor.NewBaseAdaptor()

	if lo == nil {
		lo = logrus.New()
	}
	l.config(opts, lo)
	l.entry = logrus.NewEntry(lo)

	l.SetLevel(log.DebugLevel)
	return l
}

func (l *logger) config(opts *options, lo *logrus.Logger) {
	if opts.out != nil {
		lo.SetOutput(opts.out)
	}
	if opts.formatter != nil {
		lo.SetFormatter(opts.formatter)
	} else {
		defaultFormatter := logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
		}
		lo.SetFormatter(&defaultFormatter)
	}

	if opts.include != nil {
		lo.SetReportCaller(*opts.include)
	}

	if opts.hook != nil {
		for _, hook := range opts.hook {
			lo.AddHook(hook)
		}
	}

	if opts.t != nil {
		lo.WithTime(*opts.t)
	}
}

type builder struct {
	logger *logrus.Logger
}

func NewBuilder(logger *logrus.Logger) log.Builder {
	b := new(builder)
	b.logger = logger
	return b
}

func (b *builder) Build(opts ...log.BuildOption) log.Logger {
	defaultOpts := &options{}
	for _, o := range opts {
		o(defaultOpts)
	}
	return newLogger(defaultOpts, b.logger)
}
