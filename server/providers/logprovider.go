package providers

type LogLevel int // logger level type

const (
	LOG_LEVEL_DEBUG    LogLevel = iota
	LOG_LEVEL_INFO     LogLevel = iota
	LOG_LEVEL_WARN     LogLevel = iota
	LOG_LEVEL_ERROR    LogLevel = iota
	LOG_LEVEL_CRITICAl LogLevel = iota
)

func (l LogLevel) String() string {
	switch l {
	case LOG_LEVEL_DEBUG:
		return "DEBUG"
	case LOG_LEVEL_INFO:
		return "INFO"
	case LOG_LEVEL_WARN:
		return "WARN"
	case LOG_LEVEL_ERROR:
		return "ERROR"
	case LOG_LEVEL_CRITICAl:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

type LogProvider interface {
	Flush()
	Log(level LogLevel, message string)
	Logf(level LogLevel, format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}
