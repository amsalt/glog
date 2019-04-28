package test

import (
	"testing"

	"github.com/amsalt/glog"
	"github.com/amsalt/glog/adaptor/logrus"
	"github.com/amsalt/glog/adaptor/zaplog"
	"go.uber.org/zap"
)

func TestLogrus(t *testing.T) {
	logger := logrus.NewBuilder(nil).Build()
	glog.SetLogger(logger)

	glog.Debugf("this is a test logrus log at level: %+v", glog.GetLevel())
}

func TestZap(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	l := zaplog.NewBuilder(sugar).Build()
	glog.SetLogger(l)

	glog.Debugf("this is a test zap log at level: %+v", glog.GetLevel())
}
