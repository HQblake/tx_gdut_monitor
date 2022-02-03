package output

import (
	"fmt"
	"strings"
)

type Level uint32

// 错误等级
const (
	InfoLevel Level = iota
	WarnLevel
	ErrorLevel
	PanicLevel
)
func ParseLevel(lvl string) Level {
	switch strings.ToLower(lvl) {
	case "panic":
		return PanicLevel
	case "error":
		return ErrorLevel
	case "warn", "warning":
		return WarnLevel
	case "info":
		return InfoLevel
	}
	return InfoLevel
}

func (level Level) String() string {
	if b, err := level.MarshalText(); err == nil {
		return string(b)
	} else {
		return "unknown"
	}
}
func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case WarnLevel:
		return []byte("warning"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case InfoLevel:
		return []byte("info"), nil
	case PanicLevel:
		return []byte("panic"), nil
	}

	return nil, fmt.Errorf("not a valid logrus Level %d", level)
}