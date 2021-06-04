package logger

import (
	"fmt"
	"log"
	"strings"
)

type simpleLoggerImpl struct {
}

func trivialLogger(v ...interface{}) {
	if len(v) > 0 {
		switch t := v[0].(type) {
		case string:
			if len(v) > 1 {
				fmt.Printf(t, v[1:]...)
			} else {
				fmt.Print(v...)
			}
		default:
			fmt.Print(v...)
		}
	}
	fmt.Println()
}

func (simpleLoggerImpl) Success(caller *CallerInfo, v ...interface{}) {
	trivialLogger(v...)
}

func (simpleLoggerImpl) Failed(caller *CallerInfo, v ...interface{}) {
	trivialLogger(v...)
}

func (simpleLoggerImpl) Fatal(caller *CallerInfo, v ...interface{}) {
	trivialLogger(v...)
	log.Fatal("fatal error")
}

func (simpleLoggerImpl) Error(caller *CallerInfo, v ...interface{}) {
	trivialLogger(v...)
}

func (simpleLoggerImpl) Warning(caller *CallerInfo, v ...interface{}) {
	trivialLogger(v...)
}

func (simpleLoggerImpl) Notice(caller *CallerInfo, v ...interface{}) {
	trivialLogger(v...)
}

func (simpleLoggerImpl) Debug(caller *CallerInfo, v ...interface{}) {
	trivialLogger(v...)
}

var defaultLogger GoLogger
var logLevel int

func init() {
	defaultLogger = new(simpleLoggerImpl)
	logLevel = LogNotice
}

func Debug(v ...interface{}) {
	if logLevel <= LogDebug {
		defaultLogger.Debug(GetCallerInfo(1), v...)
	}
}

func Notice(v ...interface{}) {
	if logLevel <= LogNotice {
		defaultLogger.Notice(GetCallerInfo(1), v...)
	}
}

func Success(v ...interface{}) {
	if logLevel <= LogSuccess {
		defaultLogger.Success(GetCallerInfo(1), v...)
	}
}

func Failed(v ...interface{}) {
	if logLevel <= LogFailed {
		defaultLogger.Failed(GetCallerInfo(1), v...)
	}
}

func Warning(v ...interface{}) {
	if logLevel <= LogWarning {
		defaultLogger.Warning(GetCallerInfo(1), v...)
	}
}

func Error(v ...interface{}) {
	if logLevel <= LogError {
		defaultLogger.Error(GetCallerInfo(1), v...)
	}
}

func Fatal(v ...interface{}) {
	defaultLogger.Fatal(GetCallerInfo(1), v...)
}

func SetLogger(newlogger GoLogger) {
	defaultLogger = newlogger
}

func SetLevel(level string) {
	if level != `` {
		if val, ok := logLevelLookup[strings.ToLower(level)]; ok {
			logLevel = val
		} else {
			Fatal("LogLevel '%s' not exists", level)
		}
	}
}

func GetLevel() string {
	for k, v := range logLevelLookup {
		if logLevel == v {
			return k
		}
	}
	return "none"
}
