package golog

// Logger core logger interface
type Logger interface {
	Log(log Log)
	WithField(key string, value interface{}) Entry
	WithFields(fields F) Entry
	WithError(err error) Entry
	WithName(name string) Entry
	Print(args ...any)
	Printf(format string, args ...any)
	Info(args ...any)
	Infon(name string, args ...any)
	Infof(format string, args ...any)
	Debug(args ...any)
	Debugn(name string, args ...any)
	Debugf(format string, args ...any)
	DebugJson(data interface{})
	Warn(args ...any)
	Warnn(name string, args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorn(name string, args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fataln(name string, args ...any)
	Fatalf(format string, args ...any)
}

// Entry core child logger interface
type Entry interface {
	Log(log Log)
	WithField(key string, value interface{}) Entry
	WithFields(fields F) Entry
	WithError(err error) Entry
	WithName(name string) Entry
	Print(args ...any)
	Printf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Debug(args ...any)
	Debugf(format string, args ...any)
	DebugJson(data interface{})
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

// Adapter core adapter interface
type Adapter interface {
	Log(log Log)
	Format(log *Log)
}
