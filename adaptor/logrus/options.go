package logrus

import (
	"io"
	"time"

	"github.com/amsalt/log"
	"github.com/sirupsen/logrus"
)

type options struct {
	out       io.Writer
	formatter logrus.Formatter
	include   *bool
	hook      []logrus.Hook
	t         *time.Time
}

// SetOutput sets the standard logger output.
func WithOutput(out io.Writer) log.BuildOption {
	return func(o interface{}) {
		o.(*options).out = out
	}
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) log.BuildOption {
	return func(o interface{}) {
		o.(*options).formatter = formatter
	}
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include *bool) log.BuildOption {
	return func(o interface{}) {
		o.(*options).include = include
	}
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook []logrus.Hook) log.BuildOption {
	return func(o interface{}) {
		o.(*options).hook = hook
	}
}

// WithTime creats an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t *time.Time) log.BuildOption {
	return func(o interface{}) {
		o.(*options).t = t
	}
}
