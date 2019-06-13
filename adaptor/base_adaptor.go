package adaptor

import "github.com/amsalt/log"

type BaseAdaptor struct {
	level            log.Level
	printCallerLevel log.Level
}

func NewBaseAdaptor() *BaseAdaptor {
	ba := new(BaseAdaptor)
	ba.SetLevel(log.DebugLevel)
	ba.SetPrintCallerLevel(log.DebugLevel)
	return ba
}

func (l *BaseAdaptor) SetLevel(lv log.Level) {
	l.level = lv
}

func (l *BaseAdaptor) GetLevel() log.Level {
	return l.level
}

func (l *BaseAdaptor) SetPrintCallerLevel(lv log.Level) {
	l.printCallerLevel = lv
}

func (l *BaseAdaptor) GetPrintCallerLevel() log.Level {
	return l.printCallerLevel
}

func (l *BaseAdaptor) IsEnabledLevel(lv log.Level) bool {
	return l.level <= lv
}
