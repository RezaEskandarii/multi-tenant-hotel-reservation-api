package applogger

import (
	"encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	infoLevel    = "info"
	warningLevel = "warning"
	debugLevel   = "debug"
	errorLevel   = "error"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

// getEncoder returns zapcore encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter returns zapcore WriteSyncer
func getLogWriter() zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   getLogFileName(),
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func writeLog(level string, message string) {

	InitLogger()
	defer sugarLogger.Sync()

	switch strings.TrimSpace(level) {
	case "errorLevel":
		sugarLogger.Error(message)
		break
	case "infoLevel":
		sugarLogger.Info(message)
		break
	case "warningLevel":
		sugarLogger.Warn(message)
		break
	case "debugLevel":
		sugarLogger.Debug(message)
		break
	default:
		sugarLogger.Error(message)
		break
	}
}

// LogInfo Logs in Info level.
func (l *AppLogger) LogInfo(message string) {

	writeLog(infoLevel, message)
}

// LogWarning Logs in Warning level.
func (l *AppLogger) LogWarning(message string) {

	writeLog(warningLevel, message)
}

// LogDebug Logs in DebugMode level.
func (l *AppLogger) LogDebug(message interface{}) {

	writeLog(debugLevel, fmt.Sprintf("%s", message))
}

// LogError Logs in Error level.
func (l *AppLogger) LogError(err interface{}) {

	writeLog(errorLevel, fmt.Sprintf("%s", err))
}

// LogInfoJSON store log with json pattern.
func (l *AppLogger) LogInfoJSON(message interface{}) {

	result, _ := json.Marshal(message)
	fmt.Println(string(result))
	writeLog(infoLevel, string(result))
}

func getLogFileName() string {

	fileName := "./logs/application.log"
	dir, err := os.Getwd()

	if err == nil {
		return fmt.Sprintf("%s/%s", dir, "logs/application.log")
	}

	return fileName
}