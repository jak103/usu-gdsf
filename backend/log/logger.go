package log

import (
	"fmt"
	"log"
	"runtime"
)

var logger Logger

type Logger struct {
	err error
}

func (l Logger) Error(message string, v ...any) {
	l.printMessage("[ERROR]", message, v...)
}

func (l Logger) Warn(message string, v ...any) {
	l.printMessage("[WARN] ", message, v...)
}

func (l Logger) Info(message string, v ...any) {
	l.printMessage("[INFO] ", message, v...)
}

func (l Logger) Debug(message string, v ...any) {
	l.printMessage("[DEBUG]", message, v...)
}

func (l *Logger) printMessage(level, message string, v ...any) {
	skip := 3
	if l.err != nil {
		skip--
	}

	_, filename, line, _ := runtime.Caller(skip)
	location := fmt.Sprintf("[%s:%d]", filename, line)

	message = fmt.Sprintf(level+" "+location+" "+message, v...)

	if l.err != nil {
		message += fmt.Sprintf(" :: error=%T, message=%s", l.err, l.err.Error())
	}
	log.Println(message)
	l.err = nil
}

//----------------------------

func WithError(err error) Logger {
	logger.err = err
	return logger
}

func Error(message string, v ...any) {
	logger.Error(message, v...)
}

func Warn(message string, v ...any) {
	logger.Warn(message, v...)
}

func Info(message string, v ...any) {
	logger.Info(message, v...)
}

func Debug(message string, v ...any) {
	logger.Debug(message, v...)
}
