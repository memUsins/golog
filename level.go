package glog

// Level core level declaration
type Level int

const (
	DebugLevel Level = iota - 2
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	UnselectedLevel
)

// IsEnabled check is level enabled
func (l Level) IsEnabled(level Level) bool {
	return level >= l
}

// String returns the string representation of the level
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case UnselectedLevel:
		return "unselected"
	default:
		return ""
	}
}
