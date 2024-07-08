package log

import (
	"fmt"
	"io"
)

type Logger interface {
	Debug(message string)
	Debugf(format string, args ...interface{})
	Info(message string)
	Infof(format string, args ...interface{})
	Warn(message string)
	Warnf(format string, args ...interface{})
	Error(message string)
	Errorf(format string, args ...interface{})
}

type logger struct {
	writer io.Writer
	level  int
}

const (
	LevelDebug int = iota
	LevelInfo
	LevelWarn
	LevelError
)

func NewLogger(level int, w io.Writer) Logger {
	return &logger{
		writer: w,
		level:  level,
	}
}

func (l *logger) Debug(msg string) {
	if l.level >= LevelDebug {
		l.writer.Write([]byte(msg))
	}
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *logger) Info(msg string) {
	if l.level >= LevelInfo {
		l.writer.Write([]byte(msg))
	}
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *logger) Warn(msg string) {
	if l.level >= LevelWarn {
		l.writer.Write([]byte(msg))
	}
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *logger) Error(msg string) {
	if l.level >= LevelError {
		l.writer.Write([]byte(msg))
	}
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}
