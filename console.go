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

func (c *consoleLogger) Info(message string) {
	c.logInner(INFO, message)
}

func (c *consoleLogger) Debug(message string) {
	c.logInner(DEBUG, message)
}

func (c *consoleLogger) Warning(message string) {
	c.logInner(WARNING, message)
}

func (c *consoleLogger) Error(message string) {
	c.logInner(ERROR, message)
}

func (c *consoleLogger) logInner(level logLevel, message string) {
	if c.state {
		fmt.Print(formatLogString(level, c.owner, message))
	}
}

func (c *consoleLogger) SetWriteState(state bool) {
	c.state = state
}

func (c *consoleLogger) Dispose() {
	// disposed lol
}
