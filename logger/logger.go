package logger

import (
	log "log"
)

func Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Warnf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
