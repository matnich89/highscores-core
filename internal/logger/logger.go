package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func New() *Logger {
	return &Logger{
		infoLog:  log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "ERR ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) OutputInfo(text string) {
	l.infoLog.Println(text)
}

func (l *Logger) OutputFatal(text string) {
	l.errorLog.Fatal(text)
}
