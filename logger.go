package raft

import "github.com/sirupsen/logrus"

type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fields() Fields
	Prefix() string
	WithFields(fields Fields) Logger
	SetLevel(level Level)
}

type RaftLogger struct {
	log *logrus.Logger

	prefix string
	fields Fields
}

type Fields map[string]interface{}
type Level logrus.Level

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

var _ Logger = (*RaftLogger)(nil)

func NewDefaultLogger() Logger {
	var l = logrus.New()

	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	return NewLogger(l, "RaftK", nil)
}

func NewLogger(l *logrus.Logger, perfix string, fields Fields) Logger {
	return &RaftLogger{
		log:    l,
		prefix: perfix,
		fields: fields,
	}
}

// Debug implements Logger
func (l *RaftLogger) Debug(args ...interface{}) {
	l.Log().Debug(args...)
}

// Debugf implements Logger
func (l *RaftLogger) Debugf(format string, args ...interface{}) {
	l.Log().Debugf(format, args...)
}

// Error implements Logger
func (l *RaftLogger) Error(args ...interface{}) {
	l.Log().Error(args...)
}

// Errorf implements Logger
func (l *RaftLogger) Errorf(format string, args ...interface{}) {
	l.Log().Errorf(format, args...)
}

// Fatal implements Logger
func (l *RaftLogger) Fatal(args ...interface{}) {
	l.Log().Fatal(args...)
}

// Fatalf implements Logger
func (l *RaftLogger) Fatalf(format string, args ...interface{}) {
	l.Log().Fatalf(format, args...)
}

// Info implements Logger
func (l *RaftLogger) Info(args ...interface{}) {
	l.Log().Info(args...)
}

// Infof implements Logger
func (l *RaftLogger) Infof(format string, args ...interface{}) {
	l.Log().Infof(format, args...)
}

// Panic implements Logger
func (l *RaftLogger) Panic(args ...interface{}) {
	l.Log().Panic(args...)
}

// Panicf implements Logger
func (l *RaftLogger) Panicf(format string, args ...interface{}) {
	l.Log().Panicf(format, args...)
}

// Print implements Logger
func (l *RaftLogger) Print(args ...interface{}) {
	l.Log().Print(args...)
}

// Printf implements Logger
func (l *RaftLogger) Printf(format string, args ...interface{}) {
	l.Log().Printf(format, args...)
}

// SetLevel implements Logger
func (l *RaftLogger) SetLevel(level Level) {
	l.log.SetLevel(logrus.Level(level))
}

// Trace implements Logger
func (l *RaftLogger) Trace(args ...interface{}) {
	l.Log().Trace(args...)
}

// Tracef implements Logger
func (l *RaftLogger) Tracef(format string, args ...interface{}) {
	l.Log().Tracef(format, args...)
}

// Warn implements Logger
func (l *RaftLogger) Warn(args ...interface{}) {
	l.Log().Warn(args...)
}

// Warnf implements Logger
func (l *RaftLogger) Warnf(format string, args ...interface{}) {
	l.Log().Warnf(format, args...)
}

// WithFields implements Logger
func (l *RaftLogger) WithFields(fields Fields) Logger {
	return NewLogger(l.log, l.Prefix(), l.Fields())
}

// Fields implements Logger
func (l *RaftLogger) Fields() Fields {
	return l.fields
}

// Prefix implements Logger
func (l *RaftLogger) Prefix() string {
	return l.prefix
}

func (l *RaftLogger) Log() *logrus.Entry {
	return l.log.WithField("perfix", l.prefix).WithFields(logrus.Fields(l.fields))
}
