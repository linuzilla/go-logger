package logger

func init() {
	LogLevelMap = make(map[int]string)
	for k, v := range logLevelLookup {
		LogLevelMap[v] = k
	}
}
