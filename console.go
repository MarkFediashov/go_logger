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

func (c *consoleLogger) Log(message string) {
	if c.state {
		fmt.Print(formatLogString(c.owner, message))
	}
}

func (c *consoleLogger) SetWriteState(state bool) {
	c.state = state
}

func (c *consoleLogger) Dispose() {
	// disposed lol
}
