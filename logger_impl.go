package logger

import (
	"bytes"
	"fmt"
	"log"
)

type justLogger struct {
}

func (logger *justLogger) customerLogger(bufferPtr *bytes.Buffer, v ...interface{}) string {
	var tag string

	fmt.Fprintf(bufferPtr, ">> ")
	if len(v) > 0 {
		switch t := v[0].(type) {
		//case *errcode.ErrorCode:
		//	fmt.Fprint(bufferPtr, t.ToString())
		//case errcode.ErrorCode:
		//	fmt.Fprint(bufferPtr, t.ToString())
		case string:
			if len(v) > 1 {
				fmt.Fprintf(bufferPtr, t, v[1:]...)
			} else {
				fmt.Fprint(bufferPtr, t)
			}
		default:
			if len(v) == 1 {
				fmt.Fprintf(bufferPtr, "(%T) %v", v[0], v[0])
			} else {
				fmt.Fprint(bufferPtr, v...)
			}
		}
	}
	return tag
}

func (logger *justLogger) customerLogWriter(caller *CallerInfo, logtype int, data string) {
	var buffer bytes.Buffer

	//fmt.Fprintf(&buffer, "[%s] ", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Fprint(&buffer, data)
	fmt.Fprint(&buffer, ` - `)

	if level, found := LogLevelMap[logtype]; found {
		fmt.Fprintf(&buffer, "<%s>", level)
	} else {
		fmt.Fprint(&buffer, "<unknown>")
	}

	fmt.Fprintf(&buffer, " ... [ at %s:%d ]\n", caller.FileName(), caller.LineNumber())

	fmt.Print(buffer.String())
}

func (logger *justLogger) logging(caller *CallerInfo, logtype int, v ...interface{}) {
	var buffer bytes.Buffer

	var vv []interface{}

	for i, item := range v {
		// utils.TracePrint(0, 5, utils.Detail(item))
		switch item.(type) {
		case error:
			logger.customerLogger(&buffer, v[i:]...)
			logger.customerLogWriter(caller, logtype, buffer.String())
			return
		case string:
			logger.customerLogger(&buffer, v[i:]...)
			logger.customerLogWriter(caller, logtype, buffer.String())
			return
		default:
			vv = append(vv, item)
		}
	}

	logger.customerLogger(&buffer, vv...)
	logger.customerLogWriter(caller, logtype, buffer.String())
}

func (logger *justLogger) Debug(caller *CallerInfo, v ...interface{}) {
	logger.logging(caller, LogDebug, v...)
}

func (logger *justLogger) Error(caller *CallerInfo, v ...interface{}) {
	logger.logging(caller, LogError, v...)
}

func (logger *justLogger) Warning(caller *CallerInfo, v ...interface{}) {
	logger.logging(caller, LogWarning, v...)
}

func (logger *justLogger) Notice(caller *CallerInfo, v ...interface{}) {
	logger.logging(caller, LogNotice, v...)
}

func (logger *justLogger) Success(caller *CallerInfo, v ...interface{}) {
	logger.logging(caller, LogSuccess, v...)
}

func (logger *justLogger) Failed(caller *CallerInfo, v ...interface{}) {
	logger.logging(caller, LogFailed, v...)
}

func (logger *justLogger) Fatal(caller *CallerInfo, v ...interface{}) {
	logger.logging(caller, LogFatal, v...)
	log.Fatal("fatal error")
}

func New() GoLogger {
	return &justLogger{}
}
