package logger

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	logger := &lumberjack.Logger{
		Filename:  "app.log",
		MaxSize:   20, // MB
		MaxAge:    28, // days
		Compress:  true,
		LocalTime: false,
	}

	multiWriter := io.MultiWriter(os.Stdout, logger)
	log.SetOutput(multiWriter)
}
