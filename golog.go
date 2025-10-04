package golog

import "encoding/json"

// logger implements Logger
type logger struct {
	adapters []Adapter
}

// Log core method to logging
func (l *logger) Log(log Log) {
	for _, a := range l.adapters {
		c := NewLogCopy(log)
		a.Log(c)
	}
}

// WithField returns Entry with field
func (l *logger) WithField(key string, value interface{}) Entry {
	e := NewEntry(l)
	e.WithField(key, value)
	return e
}

// WithFields returns Entry with fields
func (l *logger) WithFields(fields F) Entry {
	e := NewEntry(l)
	e.WithFields(fields)
	return e
}

// WithName returns Entry with name
func (l *logger) WithName(name string) Entry {
	e := NewEntry(l)
	e.WithName(name)
	return e
}

// WithError returns Entry with error
func (l *logger) WithError(err error) Entry {
	e := NewEntry(l)
	e.WithError(err)
	return e
}

// Print default log
func (l *logger) Print(args ...any) {
	l.Log(NewDefaultLog(UnselectedLevel, args...))
}

// Printf log with format
func (l *logger) Printf(format string, args ...any) {
	l.Log(NewDefaultLogf(UnselectedLevel, format, args...))
}

// Info default log
func (l *logger) Info(args ...any) {
	l.Log(NewDefaultLog(InfoLevel, args...))
}

// Infon log with name
func (l *logger) Infon(name string, args ...any) {
	l.Log(NewDefaultLogn(InfoLevel, name, args...))
}

// Infof log with format
func (l *logger) Infof(format string, args ...any) {
	l.Log(NewDefaultLogf(InfoLevel, format, args...))
}

// Debug default log
func (l *logger) Debug(args ...any) {
	l.Log(NewDefaultLog(DebugLevel, args...))
}

// Debugn log with name
func (l *logger) Debugn(name string, args ...any) {
	l.Log(NewDefaultLogn(DebugLevel, name, args...))
}

// Debugf log with format
func (l *logger) Debugf(format string, args ...any) {
	l.Log(NewDefaultLogf(DebugLevel, format, args...))
}

// DebugJson log formating json struct
func (l *logger) DebugJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		l.Log(NewDefaultLog(DebugLevel, "failed to unmarshal json"))
		return
	}

	l.Log(NewDefaultLog(DebugLevel, string(jsonData)))
}

// Warn default log
func (l *logger) Warn(args ...any) {
	l.Log(NewDefaultLog(WarnLevel, args...))
}

// Warnn log with name
func (l *logger) Warnn(name string, args ...any) {
	l.Log(NewDefaultLogn(WarnLevel, name, args...))
}

// Warnf log with format
func (l *logger) Warnf(format string, args ...any) {
	l.Log(NewDefaultLogf(WarnLevel, format, args...))
}

// Error default log
func (l *logger) Error(args ...any) {
	l.Log(NewDefaultLog(ErrorLevel, args...))
}

// Errorn log with name
func (l *logger) Errorn(name string, args ...any) {
	l.Log(NewDefaultLogn(ErrorLevel, name, args...))
}

// Errorf log with format
func (l *logger) Errorf(format string, args ...any) {
	l.Log(NewDefaultLogf(ErrorLevel, format, args...))
}

// Fatal default log
func (l *logger) Fatal(args ...any) {
	l.Log(NewDefaultLog(FatalLevel, args...))
}

// Fataln log with name
func (l *logger) Fataln(name string, args ...any) {
	l.Log(NewDefaultLogn(FatalLevel, name, args...))
}

// Fatalf log with format
func (l *logger) Fatalf(format string, args ...any) {
	l.Log(NewDefaultLogf(FatalLevel, format, args...))
}

func NewLogger(adapters ...Adapter) Logger {
	return &logger{adapters}
}
