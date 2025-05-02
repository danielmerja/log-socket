package log

import "time"

const (
	LTrace Level = iota
	LDebug
	LInfo
	LNotice
	LWarn
	LError
	LPanic
	LFatal
)

type (
	LogWriter chan Entry
	Level     int

	Client struct {
		LogLevel    Level `json:"level"`
		writer      LogWriter
		initialized bool
	}
	Entry struct {
		Timestamp time.Time `json:"timestamp"`
		Output    string    `json:"output"`
		File      string    `json:"file"`
		Level     string    `json:"level"`
		level     Level
	}
	Logger struct {
		FileInfoDepth int
	}
)
