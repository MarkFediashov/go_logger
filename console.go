package logger

import (
	"fmt"
	"time"
)

func NewConsoleLogger(owner string) Logger {
	return &consoleLogger{loggerState: loggerState{state: true, owner: formattedName(owner)}}
}

type consoleLogger struct {
	loggerState
}

func (c *consoleLogger) Log(data ...any) {
	if c.state {
		fmt.Printf(format, time.Now().Format("2006-01-02T15:04:05.999"), c.owner, data)
	}
}

func (c *consoleLogger) SetWriteState(state bool) {
	c.state = state
}

func (c *consoleLogger) Dispose() {
	// disposed lol
}
