package log

import "log"

func Error(message string, v ...any) {
	message = "[ERROR] " + message
	log.Printf(message, v...)
}

func Warn(message string, v ...any) {
	message = "[WARN] " + message
	log.Printf(message, v...)
}

func Info(message string, v ...any) {
	message = "[INFO] " + message
	log.Printf(message, v...)
}

func Debug(message string, v ...any) {
	message = "[DEBUG] " + message
	log.Printf(message, v...)
}
