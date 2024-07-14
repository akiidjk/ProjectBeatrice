package logger

import (
	"backend/utils"
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
		debugLogger: log.New(os.Stdout, utils.Gray+"DEBUG: "+utils.White, log.LstdFlags),
		infoLogger:  log.New(os.Stdout, utils.Cyan+"INFO: "+utils.White, log.LstdFlags),
		succeLogger: log.New(os.Stdout, utils.Green+"SUCCESS: "+utils.White, log.LstdFlags),
		warnLogger:  log.New(os.Stdout, utils.Yellow+"WARN: "+utils.White, log.LstdFlags),
		errorLogger: log.New(os.Stdout, utils.Red+"ERROR: "+utils.White, log.LstdFlags),
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
