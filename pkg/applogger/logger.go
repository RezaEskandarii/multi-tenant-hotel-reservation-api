package applogger

import (
	"github.com/amoghe/distillog"
	"github.com/natefinch/lumberjack"
	"strings"
)

type Logger interface {
	LogInfo(message interface{})
	LogDebug(message interface{})
	LogWarning(message interface{})
	LogError(message interface{})
}

type AppLogger struct {
}

func New() *AppLogger {

	return &AppLogger{}
}

var (
	logHandler = &lumberjack.Logger{
		Filename:   "logs/application.log",
		MaxSize:    5, // megabytes
		MaxBackups: 10,
		MaxAge:     30, // days
		Compress:   false,
	}
)

const (
	infoLevel    = "info"
	warningLevel = "warning"
	debugLevel   = "debug"
	errorLevel   = "error"
)

func writeLog(level string, message interface{}) {

	logger := distillog.NewStreamLogger(level, logHandler)
	defer logger.Close()
	distillog.SetOutput(logHandler)

	switch strings.TrimSpace(level) {
	case "errorLevel":
		logger.Errorf("%s", message)
		break
	case "infoLevel":
		logger.Infof("%s", message)
		break
	case "warningLevel":
		logger.Warningf("%s", message)
		break
	case "debugLevel":
		logger.Debugf("%s", message)
		break
	default:
		logger.Errorf("%s", message)
		break
	}
}

// LogInfo Logs in Info level.
func (l *AppLogger) LogInfo(message interface{}) {

	writeLog(infoLevel, message)
}

// LogWarning Logs in Warning level.
func (l *AppLogger) LogWarning(message interface{}) {

	writeLog(warningLevel, message)
}

// LogDebug Logs in Debug level.
func (l *AppLogger) LogDebug(message interface{}) {

	writeLog(debugLevel, message)
}

// LogError Logs in Error level.
func (l *AppLogger) LogError(message interface{}) {

	writeLog(errorLevel, message)
}
