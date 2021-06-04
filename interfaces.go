package logger

type GoLogger interface {
	Debug(caller *CallerInfo, v ...interface{})
	Notice(caller *CallerInfo, v ...interface{})
	Warning(caller *CallerInfo, v ...interface{})
	Error(caller *CallerInfo, v ...interface{})
	Success(caller *CallerInfo, v ...interface{})
	Failed(caller *CallerInfo, v ...interface{})
	Fatal(caller *CallerInfo, v ...interface{})
}
