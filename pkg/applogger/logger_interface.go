package applogger

import (
	"gorm.io/gorm"
	"time"
)

type Log struct {
	gorm.Model
	Message string
	Level   string
	UtcTime time.Time
}

type LogDbWriter struct {
	db *gorm.DB
}

func NewLogDbWriter(db *gorm.DB) *LogDbWriter {
	db.AutoMigrate(&Log{})
	return &LogDbWriter{db: db}
}

func (w *LogDbWriter) Write(p []byte) (n int, err error) {
	log := &Log{
		Message: string(p),
		Level:   "info",
	}
	log.CreatedAt = time.Now()
	log.UtcTime = time.Now().UTC()
	if err := w.db.Create(log).Error; err != nil {
		return 0, err
	}
	return len(p), nil
}

// Logger is logger base type
type Logger interface {
	LogInfo(message string)
	LogInfoJSON(message interface{})
	LogDebug(message interface{})
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

// New returns AppLogger pointer
func New(conf *LoggerConfig) *AppLogger {

	return &AppLogger{
		Config: conf,
	}
}
