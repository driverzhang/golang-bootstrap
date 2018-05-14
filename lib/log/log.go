package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"fmt"
	"github.com/go-stack/stack"
)

var logger *logrus.Logger
var debug = false

// debug: 使用text格式, Level是Debug, 打印所有级别
// not debug: 使用json格式, level是Info, 不打印Debug级别
func SetDebug(d bool) {
	debug = d
	if debug {
		format := new(logrus.TextFormatter)
		//format.ForceColors = true
		format.TimestampFormat = "2006-01-02 15:04:05"
		logger.Level = logrus.DebugLevel
		logger.Formatter = format
	} else {
		format := new(logrus.JSONFormatter)
		format.TimestampFormat = "2006-01-02 15:04:05"
		logger.Level = logrus.InfoLevel
		logger.Formatter = format
	}
}

func WithField(key string, value interface{}) *logrus.Entry {
	return withCaller().WithField(key, value)
}

func WithFields(fs logrus.Fields) *logrus.Entry {
	return withCaller().WithFields(fs)
}

func withCaller() *logrus.Entry {
	var key = "caller"
	var value interface{}
	if debug {
		// 支持goland点击跳转
		value = fmt.Sprintf(" %+v:", stack.Caller(2))
	} else {
		value = fmt.Sprintf("%+v", stack.Caller(2))
	}

	return logger.WithFields(logrus.Fields{key: value})
}

func Error(args ...interface{}) {
	withCaller().Error(args...)
}

func Info(args ...interface{}) {
	withCaller().Info(args...)
}

func Print(args ...interface{}) {
	withCaller().Print(args...)
}

func Warn(args ...interface{}) {
	withCaller().Warn(args...)
}

func Debug(args ...interface{}) {
	withCaller().Debug(args...)
}

func Errorf(format string, args ...interface{}) {
	withCaller().Errorf(format, args...)
}

func Infof(format string, args ...interface{}) {
	withCaller().Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	withCaller().Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	withCaller().Warnf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	withCaller().Debugf(format, args...)
}

func init() {
	logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: nil,
		Hooks:     make(logrus.LevelHooks),
		Level:     0,
	}

	SetDebug(true)
}
