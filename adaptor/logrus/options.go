package logrus

import (
	"io"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/amsalt/glog"
)

type options struct {
	out       io.Writer
	formatter logrus.Formatter
	include   *bool
	hook      []logrus.Hook
	t         *time.Time
}

// SetOutput sets the standard logger output.
func WithOutput(out io.Writer) glog.BuildOption {
	return func(o interface{}) {
		o.(*options).out = out
	}
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) glog.BuildOption {
	return func(o interface{}) {
		o.(*options).formatter = formatter
	}
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include *bool) glog.BuildOption {
	return func(o interface{}) {
		o.(*options).include = include
	}
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook []logrus.Hook) glog.BuildOption {
	return func(o interface{}) {
		o.(*options).hook = hook
	}
}

// WithTime creats an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t *time.Time) glog.BuildOption {
	return func(o interface{}) {
		o.(*options).t = t
	}
}
