package glog

import (
	"fmt"
	"time"
)

// F alias for map
type F map[string]interface{}

// LogData core log props interface
type LogData struct {
	Fields   F
	Error    error
	Name     string
	WithName bool
}

// Log core log interface
type Log struct {
	Level     Level
	Data      LogData
	Message   string
	Timestamp time.Time
}

// NewLogCopy returns copied Log
func NewLogCopy(log Log) Log {
	fields := make(map[string]interface{})
	for k, v := range log.Data.Fields {
		fields[k] = v
	}

	return Log{
		Level:   log.Level,
		Message: log.Message,
		Data: LogData{
			Fields:   fields,
			Error:    log.Data.Error,
			Name:     log.Data.Name,
			WithName: log.Data.WithName,
		},
		Timestamp: log.Timestamp,
	}
}

// NewDefaultLog returns new Log
func NewDefaultLog(level Level, args ...interface{}) Log {
	return Log{
		Level:   level,
		Message: fmt.Sprint(args...),
		Data: LogData{
			Fields: make(F),
		},
		Timestamp: time.Now(),
	}
}

// NewDefaultLogn returns new Log with name args
func NewDefaultLogn(level Level, name string, args ...interface{}) Log {
	return Log{
		Level:   level,
		Message: fmt.Sprint(args...),
		Data: LogData{
			Name:   name,
			Fields: make(F),
		},
		Timestamp: time.Now(),
	}
}

// NewDefaultLogf returns new Log with format args
func NewDefaultLogf(level Level, format string, args ...interface{}) Log {
	return Log{
		Level:   level,
		Message: fmt.Sprintf(format, args...),
		Data: LogData{
			Fields: make(F),
		},
		Timestamp: time.Now(),
	}
}
