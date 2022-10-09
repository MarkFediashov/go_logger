package logger

import (
	"fmt"
	"time"
)

type loggerState struct {
	owner string
	state bool
}

type logLevel string

const (
	DEBUG   logLevel = "DEBUG"
	INFO    logLevel = "INFO"
	WARNING logLevel = "WARNING"
	ERROR   logLevel = "ERROR"
)

func formatLogString(level logLevel, owner string, message string) string {
	return fmt.Sprintf(format, time.Now().Format("2006-01-02T15:04:05.999"), level, owner, message)
}

func formattedName(owner string) string {
	own := ""
	if len(owner) <= sourceFieldLength {
		own = owner
	} else if owner == "" {
		own = "Anonymous"
	} else {
		own = owner[:15]
	}

	return own
}
