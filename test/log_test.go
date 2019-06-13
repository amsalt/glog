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
	log.Debug("this is a test logrus log at level: 4")

	log.Errorf("this is a test logrus log at level: %d", 1)
	log.Error("this is a test logrus log at level: 1")

	log.SetPrintCallerLevel(log.ErrorLevel)

	log.Errorf("this is a test logrus log at level: %d", 1)
	log.Error("this is a test logrus log at level: 1")

	log.SetLevel(log.ErrorLevel)

	log.Debugf("this is a test logrus log at level: %+v", log.GetLevel())
	log.Debug("this is a test logrus log at level: 4")
}

func TestZap(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.DisableCaller = true
	logger, _ := config.Build()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	l := zaplog.NewBuilder(sugar).Build()
	log.SetLogger(l)

	log.Debugf("this is a test zap log at level: %+v", log.GetLevel())
	log.Debug("this is a test zap log at level: 4")

	log.Errorf("this is a test zap log at level: %d", 1)
	log.Error("this is a test zap log at level: 1")

	log.SetPrintCallerLevel(log.ErrorLevel)

	log.Errorf("this is a test zap log at level: %d", 1)
	log.Error("this is a test zap log at level: 1")

	log.SetLevel(log.ErrorLevel)

	log.Debugf("this is a test zap log at level: %+v", log.GetLevel())
	log.Debug("this is a test zap log at level: 4")
}
