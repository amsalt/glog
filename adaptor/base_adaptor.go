package adaptor

import "github.com/amsalt/glog"

type BaseAdaptor struct {
	level glog.Level
}

func NewBaseAdaptor() *BaseAdaptor {
	ba := new(BaseAdaptor)
	ba.SetLevel(glog.DebugLevel)
	return ba
}

func (l *BaseAdaptor) SetLevel(lv glog.Level) {
	l.level = lv
}

func (l *BaseAdaptor) GetLevel() glog.Level {
	return l.level
}

func (l *BaseAdaptor) IsEnabledLevel(lv glog.Level) bool {
	return l.level <= lv
}
