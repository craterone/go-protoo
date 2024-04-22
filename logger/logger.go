package logger

import (
	"fmt"
	log "log"
)

var EnableLog = false

func Infof(format string, v ...interface{}) {
	if EnableLog {
		format = fmt.Sprintf("[Info] %s", format)
		log.Printf(format, v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if EnableLog {
		format = fmt.Sprintf("[Debug] %s", format)
		log.Printf(format, v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if EnableLog {
		format = fmt.Sprintf("[Warn] %s", format)
		log.Printf(format, v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if EnableLog {
		format = fmt.Sprintf("[Error] %s", format)
		log.Printf(format, v...)
	}
}
