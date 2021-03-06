// Package log provides structured and customizeable logging.
package log

// Message is something that can be logged.
type Message struct {
	// The importance of the message.
	Lvl Level
	// The content of the message.
	Msg string
	// Additional data for the message.
	Fds Fields
}

// Level describes the importance of a log message.
type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	DisableLogging
)

var levelToName = map[Level]string{
	DebugLevel:     "debug",
	InfoLevel:      "info",
	WarningLevel:   "warning",
	ErrorLevel:     "error",
	DisableLogging: "disabled",
}

func (l Level) String() string {
	name, ok := levelToName[l]
	if !ok {
		return "unknown"
	}
	return name
}

// Fields describes additional data that can be attached to a log message.
type Fields map[string]interface{}

// Copy creates a copy. This is a shallow copy.
func (f Fields) Copy() Fields {
	result := make(Fields, len(f))
	for k, v := range f {
		result[k] = v
	}
	return result
}

// Update adds all entries of the given Fields to the one it is called on.
func (f Fields) Update(a Fields) {
	for k, v := range a {
		f[k] = v
	}
}
