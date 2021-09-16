package logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

//Log 暴露给外界的接口
var Log *logrus.Entry

func init() {
	logger := logrus.New()
	logger.Formatter = new(prefixed.TextFormatter)
	logger.Level = logrus.DebugLevel
	// logger.Level = logrus.InfoLevel
	Log = logger.WithFields(logrus.Fields{"prefix": "biu guess"})
}
