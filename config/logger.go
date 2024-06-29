package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info *log.Logger
	warning *log.Logger
	err *log.Logger
	writer io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug: log.New(writer, "DEBUG: ", logger.Flags()),
		info: log.New(writer, "INFO: ", logger.Flags()),
    warning: log.New(writer, "WARNING: ", logger.Flags()),
		err: log.New(writer, "ERROR: ", logger.Flags()),
    writer: writer,
	}
}

// create Non formated logs

func (l *Logger) Debug(v ... interface{}) {
	l.debug.Panicln(v ...)
}
func (l *Logger) Info(v ... interface{}) {
	l.info.Panicln(v ...)
}

func (l *Logger) Warning(v ... interface{}) {
	l.warning.Panicln(v ...)
}
func (l *Logger) Err(v ... interface{}) {
	l.err.Panicln(v ...)
}
