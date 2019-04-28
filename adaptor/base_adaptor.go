package adaptor

import "github.com/amsalt/log"

type BaseAdaptor struct {
	level log.Level
}

func NewBaseAdaptor() *BaseAdaptor {
	ba := new(BaseAdaptor)
	ba.SetLevel(log.DebugLevel)
	return ba
}

func (l *BaseAdaptor) SetLevel(lv log.Level) {
	l.level = lv
}

func (l *BaseAdaptor) GetLevel() log.Level {
	return l.level
}

func (l *BaseAdaptor) IsEnabledLevel(lv log.Level) bool {
	return l.level <= lv
}
