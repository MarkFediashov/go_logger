package logger

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(data ...any)
	SetWriteState(bool)
}

func NewLogger(owner string) Logger {
	own := ""
	if len(owner) < 15 {
		own = owner
	} else if owner == "" {
		own = "Anonymous"
	} else {
		own = owner[:15]
	}
	return &consoleLogger{state: true, owner: own}
}

type consoleLogger struct {
	state bool
	owner string
}

func (c *consoleLogger) Log(data ...any) {
	if c.state {
		fmt.Printf("|%s| [%-15s]: %s\n", time.Now().Format("2006-01-02T15:04:05.999"), c.owner, data)
	}
}

func (c *consoleLogger) SetWriteState(state bool) {
	c.state = state
}
