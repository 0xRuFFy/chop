package chop

import (
	"fmt"

	"github.com/0xruffy/gotint"
)

var ts = gotint.NewTintString("")

var (
	dimTint = gotint.NewTint().Dim()
)

type Logger struct {
	level LogLevel
}

func NewLogger() *Logger {
	return &Logger{
		level: INFO,
	}
}

func (l *Logger) Level() *LogLevel {
	return &l.level
}

// TODO: add format support
func (l *Logger) log(level LogLevel, args ...interface{}) {
	if level >= l.level {
		ts.WithTint(dimTint)

		// TODO: add timestamp / other prefixes

		fmt.Printf(
			"%s %s\t %s ",
			ts.WithContent("[").String(), // .String() so its not overwritten by the next ts.WithContent()
			level,
			ts.WithContent("]"),
		)
		for _, arg := range args {
			fmt.Print(arg)
		}
		fmt.Println()
	}
}

func (l *Logger) Trace(args ...interface{}) {
	l.log(TRACE, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.log(DEBUG, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.log(INFO, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log(WARN, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.log(ERROR, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.log(FATAL, args...)
}
