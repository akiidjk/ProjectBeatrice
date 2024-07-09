package utils

import (
	"log"
	"os"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const Gray = "\033[37m"
const White = "\033[97m"

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func Init() (*log.Logger, *log.Logger, *log.Logger) {
	InfoLogger = log.New(os.Stdout, Cyan+"INFO: "+Gray, log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, Yellow+"WARNING: "+Gray, log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, Red+"ERROR: "+Gray, log.Ldate|log.Ltime|log.Lshortfile)

	return InfoLogger, WarningLogger, ErrorLogger
}
