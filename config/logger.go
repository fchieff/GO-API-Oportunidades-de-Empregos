package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)
	return &Logger{
		debug:   log.New(writer, prefix+"[DEBUG] : ", logger.Flags()),
		info:    log.New(writer, prefix+"[INFO] : ", logger.Flags()),
		warning: log.New(writer, prefix+"[WARNING] : ", logger.Flags()),
		err:     log.New(writer, prefix+"[ERROR] : ", logger.Flags()),
		writer:  writer,
	}
}

// Create non-formatted logs
func (l *Logger) Debug(valor ...interface{}) {
	l.debug.Println(valor...)
}
func (l *Logger) Info(valor ...interface{}) {
	l.info.Println(valor...)
}
func (l *Logger) Warning(valor ...interface{}) {
	l.warning.Println(valor...)
}
func (l *Logger) Error(valor ...interface{}) {
	l.err.Println(valor...)
}

// Create formatted logs
func (l *Logger) Debugf(format string, valor ...interface{}) {
	l.debug.Printf(format, valor...)
}
func (l *Logger) Infof(format string, valor ...interface{}) {
	l.info.Printf(format, valor...)
}
func (l *Logger) Warningf(format string, valor ...interface{}) {
	l.warning.Printf(format, valor...)
}
func (l *Logger) Errorf(format string, valor ...interface{}) {
	l.err.Printf(format, valor...)
}

// Esse arquivo define um sistema de logging personalizado em Go, usando o pacote padrão log. Ele cria um Logger com quatro níveis de logs: Debug, Info, Warning e Error. Cada nível tem duas versões: uma para logs simples e outra para logs formatados.
