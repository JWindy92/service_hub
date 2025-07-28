package logging

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	gormLogger "gorm.io/gorm/logger"
)

var (
	LogWriter io.Writer
	GormLog   gormLogger.Interface
)

// Init sets up slog and GORM logging to stdout and a log file.
func InitLogger() {
	logDir := "logs"
	logPath := filepath.Join(logDir, "app.log")

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("failed to create log directory: " + err.Error())
	}

	// Open log file (append or create)
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}

	LogWriter = io.MultiWriter(os.Stdout, logFile)

	// Set slog JSON handler to write to both stdout and file
	slogHandler := slog.NewJSONHandler(LogWriter, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	slog.SetDefault(slog.New(slogHandler))

	// Set GORM logger to use same output
	GormLog = gormLogger.New(
		log.New(LogWriter, "[GORM] ", log.LstdFlags),
		gormLogger.Config{
			SlowThreshold: time.Second,
			LogLevel:      gormLogger.Info,
			Colorful:      false,
		},
	)
}

// package logging

// import (
// 	"io"
// 	"log"
// 	"log/slog"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	gormLogger "gorm.io/gorm/logger"
// )

// var (
// 	GormLogger gormLogger.Interface
// 	writer     io.Writer
// )

// func InitGormLogger() {
// 	if writer == nil {
// 		panic("LogWriter is nil — did you forget to call InitLogger() first?")
// 	}

// 	GormLogger = gormLogger.New(
// 		log.New(writer, "[GORM] ", log.LstdFlags),
// 		gormLogger.Config{
// 			SlowThreshold: time.Second,
// 			LogLevel:      gormLogger.Info,
// 			Colorful:      false,
// 		},
// 	)
// }

// func InitLogger() {
// 	logDir := "logs"
// 	logPath := filepath.Join(logDir, "app.log")

// 	// Create logs/ directory if it doesn't exist
// 	if err := os.MkdirAll(logDir, 0755); err != nil {
// 		panic("failed to create log directory: " + err.Error())
// 	}

// 	// Open log file (create if doesn't exist)
// 	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic("failed to open log file: " + err.Error())
// 	}

// 	// Write to both stdout and file
// 	multiWriter := io.MultiWriter(os.Stdout, logFile)

// 	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
// 		Level:     slog.LevelInfo,
// 		AddSource: true,
// 	})

// 	slog.SetDefault(slog.New(handler))

// 	// InitGormLogger()
// }
