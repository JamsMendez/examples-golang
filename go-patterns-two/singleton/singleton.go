package singleton

import (
	"fmt"
	"sync"
)

var logger *CustomLogger
var once sync.Once

func getLoggerInstance() *CustomLogger {
	if logger == nil {
		// creating logger instance
		logger = &CustomLogger{}
	}

	// return logger instance
	return logger
}

func getLoggerInstanceConcurrent() *CustomLogger {
	once.Do(func() {
		// creating logger instance
		logger = &CustomLogger{}
	})

	// return logger instance
	return logger
}

type CustomLogger struct {
	loglevel int
}

func (l *CustomLogger) Log(s string) {
	fmt.Println(l.loglevel, ": ", s)
}

func (l *CustomLogger) SetLogLevel(level int) {
	l.loglevel = level
}
