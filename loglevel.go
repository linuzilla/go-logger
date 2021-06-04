package logger

const (
	LogDebug = iota
	LogNotice
	LogSuccess
	LogFailed
	LogWarning
	LogError
	LogFatal
)

var logLevelLookup = map[string]int{
	"debug":   LogDebug,
	"notice":  LogNotice,
	"warning": LogWarning,
	"success": LogSuccess,
	"failed":  LogFailed,
	"error":   LogError,
	"fatal":   LogFatal,
}

var LogLevelMap map[int]string
