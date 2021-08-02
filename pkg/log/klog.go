package log

import (
	"fmt"

	"github.com/fatedier/beego/logs"
	kslog "github.com/go-kratos/kratos/v2/log"
)

type KLog struct {
	beelog *logs.BeeLogger
}

var klog *KLog
var beelog *logs.BeeLogger

var kLogHelper *kslog.Helper

func NewKLog(logWay string, logFile string, logLevel int, maxdays int64) *KLog {
	kloger := &KLog{}

	kloger.beelog = logs.NewLogger(200)
	kloger.beelog.EnableFuncCallDepth(true)
	kloger.beelog.SetLogFuncCallDepth(kloger.beelog.GetLogFuncCallDepth() + 3)

	if logWay == "console" {
		kloger.beelog.SetLogger("console", "")
	} else {
		params := fmt.Sprintf(`{"filename": "%s", "maxdays": %d}`, logFile, maxdays)
		kloger.beelog.SetLogger("file", params)
	}
	kloger.beelog.SetLevel(logLevel)

	kLogHelper = kslog.NewHelper(kloger)
	return klog
}

func (k *KLog) Log(level kslog.Level, keyvals ...interface{}) error {
	switch level {
	case kslog.LevelDebug:
		k.beelog.Debug(keyvals[1].(string))
	case kslog.LevelInfo:
		k.beelog.Info(keyvals[1].(string))
	case kslog.LevelWarn:
		k.beelog.Warn(keyvals[1].(string))
	case kslog.LevelError:
		k.beelog.Error(keyvals[1].(string))
	case kslog.LevelFatal:
		k.beelog.Critical(keyvals[1].(string))
	}

	return nil
}

func KFatal(format string, v ...interface{}) {
	kLogHelper.Fatalf(format, v...)
}

func KError(format string, v ...interface{}) {

	kLogHelper.Errorf(format, v...)
}

func KWarn(format string, v ...interface{}) {
	kLogHelper.Warnf(format, v...)
}

func KInfo(format string, v ...interface{}) {
	kLogHelper.Infof(format, v...)
}

func KDebug(format string, v ...interface{}) {
	kLogHelper.Debugf(format, v...)
}
