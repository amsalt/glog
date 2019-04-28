package test

import (
	"testing"

	"github.com/amsalt/log"
	"github.com/amsalt/log/adaptor/logrus"
	"github.com/amsalt/log/adaptor/zaplog"
	"go.uber.org/zap"
)

func TestLogrus(t *testing.T) {
	logger := logrus.NewBuilder(nil).Build()
	log.SetLogger(logger)

	log.Debugf("this is a test logrus log at level: %+v", log.GetLevel())
}

func TestZap(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	l := zaplog.NewBuilder(sugar).Build()
	log.SetLogger(l)

	log.Debugf("this is a test zap log at level: %+v", log.GetLevel())
}
