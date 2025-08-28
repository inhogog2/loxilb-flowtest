package util

import tk "github.com/loxilb-io/loxilib"

func LogString2Level(logStr string) tk.LogLevelT {
	logLevel := tk.LogDebug
	switch logStr {
	case "info":
		logLevel = tk.LogInfo
	case "error":
		logLevel = tk.LogError
	case "notice":
		logLevel = tk.LogNotice
	case "warning":
		logLevel = tk.LogWarning
	case "alert":
		logLevel = tk.LogAlert
	case "critical":
		logLevel = tk.LogCritical
	case "emergency":
		logLevel = tk.LogEmerg
	case "trace":
		logLevel = tk.LogTrace
	case "debug":
	default:
		logLevel = tk.LogDebug
	}
	return logLevel
}

func LogItInit(logFile string, logLevel tk.LogLevelT, toTTY bool) *tk.Logger {
	return tk.LogItInit(logFile, logLevel, true)
}
