package logger

import (
	"fmt"
)

func NewConsoleLogger(owner string) Logger {
	return &consoleLogger{loggerState: loggerState{state: true, owner: formattedName(owner)}}
}

type consoleLogger struct {
	loggerState
}

func (c *consoleLogger) Log(data ...any) {
	if c.state {
		fmt.Print(formatLogString(c.owner, data))
	}
}

func (c *consoleLogger) SetWriteState(state bool) {
	c.state = state
}

func (c *consoleLogger) Dispose() {
	// disposed lol
}
