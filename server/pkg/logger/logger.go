package logger

import (
	"io"
	"log"
	"os"
	"sync"
)

type LoggerConfig struct {
	WriteToConsole bool
	WriteToFile    bool
	FilePath       string
}
type CustomLogger struct {
	config *LoggerConfig
	logger *log.Logger
	mu     sync.Mutex
}

// NewCustomLogger initializes a new custom logger with time-stamped logs
func NewCustomLogger(config *LoggerConfig) (*CustomLogger, error) {
	var writers []io.Writer

	if config.WriteToConsole {
		writers = append(writers, os.Stdout)
	}

	if config.WriteToFile {
		file, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		writers = append(writers, file)
	}

	multiWriter := io.MultiWriter(writers...)
	// Set the logger to include the date and time
	logger := log.New(multiWriter, "", log.LstdFlags|log.Lmicroseconds)

	return &CustomLogger{
		config: config,
		logger: logger,
	}, nil
}


func (l *CustomLogger) Info(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Println("INFO: " ,args)
}

func (l *CustomLogger) Warn(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Println("WARN: " ,args)
}

func (l *CustomLogger) Error(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Println("ERROR: " , args)
}
func (l *CustomLogger) Fatal(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Fatal("Panic: ", args)
}

func (l *CustomLogger) Debug(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Println("Debug: " , args)
}
