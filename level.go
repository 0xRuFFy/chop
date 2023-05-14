package chop

import "github.com/0xruffy/gotint"

type LogLevel int

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

var (
	traceTint = gotint.NewTint().Bold()
	debugTint = gotint.NewTint().Bold().Blue()
	infoTint  = gotint.NewTint().Bold().Green()
	warnTint  = gotint.NewTint().Bold().Yellow()
	errorTint = gotint.NewTint().Bold().Red()
	fatalTint = gotint.NewTint().Bold().Red().OnWhite()
)

func (l LogLevel) String() string {
	ts := gotint.NewTintString("")

	switch l {
	case TRACE:
		ts.WithContent("TRACE").WithTint(traceTint)
	case DEBUG:
		ts.WithContent("DEBUG").WithTint(debugTint)
	case INFO:
		ts.WithContent("INFO").WithTint(infoTint)
	case WARN:
		ts.WithContent("WARN").WithTint(warnTint)
	case ERROR:
		ts.WithContent("ERROR").WithTint(errorTint)
	case FATAL:
		ts.WithContent("FATAL").WithTint(fatalTint)
	case OFF:
		return "OFF"
	default:
		return "UNKNOWN"
	}
	return ts.String()
}

// Set the log level to TRACE
func (ll *LogLevel) Trace() {
	*ll = TRACE
}

// Set the log level to DEBUG
func (ll *LogLevel) Debug() {
	*ll = DEBUG
}

// Set the log level to INFO
func (ll *LogLevel) Info() {
	*ll = INFO
}

// Set the log level to WARN
func (ll *LogLevel) Warn() {
	*ll = WARN
}

// Set the log level to ERROR
func (ll *LogLevel) Error() {
	*ll = ERROR
}

// Set the log level to FATAL
func (ll *LogLevel) Fatal() {
	*ll = FATAL
}

// Set the log level to OFF
func (ll *LogLevel) Off() {
	*ll = OFF
}
