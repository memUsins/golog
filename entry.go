package glog

import (
	"encoding/json"
	"sync"
)

// entry implements Entry
type entry struct {
	logger Logger
	data   LogData
	mu     sync.RWMutex
}

// Log core method to logging
func (e *entry) Log(log Log) {
	e.mu.RLock()
	log.Data = e.data
	e.mu.RUnlock()

	e.logger.Log(log)

	e.mu.Lock()
	defer e.mu.Unlock()
	e.data.Error = nil
	e.data.Fields = nil
}

// WithField add field to entry data
func (e *entry) WithField(key string, value interface{}) Entry {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.data.Fields = F{key: value}

	return e
}

// WithFields add fields to entry data
func (e *entry) WithFields(fields F) Entry {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.data.Fields = fields

	return e
}

// WithError add error to entry data
func (e *entry) WithError(err error) Entry {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.data.Error = err

	return e
}

// WithName add name to entry data
func (e *entry) WithName(name string) Entry {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.data.Name = name

	return e
}

// Print default log
func (e *entry) Print(args ...any) {
	e.Log(NewDefaultLog(UnselectedLevel, args...))
}

// Printf log with format
func (e *entry) Printf(format string, args ...any) {
	e.Log(NewDefaultLogf(UnselectedLevel, format, args...))
}

// Info default log
func (e *entry) Info(args ...any) {
	e.Log(NewDefaultLog(InfoLevel, args...))
}

// Infof log with format
func (e *entry) Infof(format string, args ...any) {
	e.Log(NewDefaultLogf(InfoLevel, format, args...))
}

// Debug default log
func (e *entry) Debug(args ...any) {
	e.Log(NewDefaultLog(DebugLevel, args...))
}

// Debugf log with format
func (e *entry) Debugf(format string, args ...any) {
	e.Log(NewDefaultLogf(DebugLevel, format, args...))
}

// DebugJson log formating json struct
func (e *entry) DebugJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		e.Log(NewDefaultLog(DebugLevel, "failed to unmarshal json"))
		return
	}

	e.Log(NewDefaultLog(DebugLevel, string(jsonData)))
}

// Warn default log
func (e *entry) Warn(args ...any) {
	e.Log(NewDefaultLog(WarnLevel, args...))
}

// Warnf log with format
func (e *entry) Warnf(format string, args ...any) {
	e.Log(NewDefaultLogf(WarnLevel, format, args...))
}

// Error default log
func (e *entry) Error(args ...any) {
	e.Log(NewDefaultLog(ErrorLevel, args...))
}

// Errorf log with format
func (e *entry) Errorf(format string, args ...any) {
	e.Log(NewDefaultLogf(ErrorLevel, format, args...))
}

// Fatal default log
func (e *entry) Fatal(args ...any) {
	e.Log(NewDefaultLog(FatalLevel, args...))
}

// Fatalf log with format
func (e *entry) Fatalf(format string, args ...any) {
	e.Log(NewDefaultLogf(FatalLevel, format, args...))
}

func NewEntry(l Logger) Entry {
	return &entry{
		logger: l,
	}
}
