package applogger

type Logger interface {
	LogInfo(message string)
	LogDebug(message string)
	LogWarning(message string)
	LogError(err interface{})
}

type LoggerConfig struct {
	CustomFilePath      string
	CustomFileExtension string
	LogServerAddr       string
}

type AppLogger struct {
	Config *LoggerConfig
}

func New(conf *LoggerConfig) *AppLogger {

	return &AppLogger{
		Config: conf,
	}
}
