package logger

import (
	"fmt"
	"time"
)

type loggerState struct {
	owner string
	state bool
}

func formatLogString(owner string, data ...any) string {
	return fmt.Sprintf(format, time.Now().Format("2006-01-02T15:04:05.999"), owner, data)
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
