package utils

import (
	"backend/config"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func logLevel() hlog.Level {
	loggerConfig := config.Config.Logger
	level := loggerConfig.Level
	switch strings.ToLower(level) {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		hlog.Warnf("Unknown log level: %s, defaulting to info", level)
		return hlog.LevelInfo
	}
}

func InitLogger() {
	hlog.SetLogger(hlog.DefaultLogger())
	hlog.SetLevel(logLevel())
}
