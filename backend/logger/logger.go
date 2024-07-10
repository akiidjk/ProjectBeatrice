package logger

import (
	"log"
	"os"
)

// Log levels
const (
	DebugLevel = iota
	InfoLevel
	SuccessLevel
	WarningLevel
	ErrorLevel
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

type Logger struct {
	Level       int
	debugLogger *log.Logger
	infoLogger  *log.Logger
	succeLogger *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

var logger *Logger

func init() {
	logger = &Logger{
		Level:       InfoLevel,
		debugLogger: log.New(os.Stdout, Gray+"DEBUG: "+White, log.LstdFlags),
		infoLogger:  log.New(os.Stdout, Cyan+"INFO: "+White, log.LstdFlags),
		succeLogger: log.New(os.Stdout, Green+"SUCCESS: "+White, log.LstdFlags),
		warnLogger:  log.New(os.Stdout, Yellow+"WARN: "+White, log.LstdFlags),
		errorLogger: log.New(os.Stdout, Red+"ERROR: "+White, log.LstdFlags),
	}
}

// Set log level
func SetLevel(level int) {
	logger.Level = level
}
func Debug(message string) {
	if logger.Level <= DebugLevel {
		logger.debugLogger.Println(message)
	}
}

func Info(message string) {
	if logger.Level <= InfoLevel {
		logger.infoLogger.Println(message)
	}
}

func Success(message string) {
	if logger.Level <= SuccessLevel {
		logger.succeLogger.Println(message)
	}
}

func Warning(message string) {
	if logger.Level <= WarningLevel {
		logger.warnLogger.Println(message)
	}
}

func Error(message string) {
	if logger.Level <= ErrorLevel {
		logger.errorLogger.Println(message)
	}
}
