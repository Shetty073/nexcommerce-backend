package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func SetupLogger() {
	Logger = log.New(os.Stdout, "nexcommerce: ", log.LstdFlags|log.Lshortfile)
}
