package loglevel

import "os"

var LogLevelDebug = "debug"
var LogLevelInfo = "info"

func LogLevel() string {
	lvl := os.Getenv("LOG_LEVEL")
	if lvl == "" {
		return "info"
	}
	return lvl
}
