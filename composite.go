package logger

import (
	"time"
)

type CompositeLogger struct {
	loggers []Logger
	console *consoleLogger
	file    *fileLogger
	network *socketLogger
}

func NewChainedLogger(owner string, fileName *string, enableConsole bool, enableFile bool, enableSocket bool, port int) (*CompositeLogger, error) {

	loggers := make([]Logger, 0, 3)

	if enableConsole {
		loggers = append(loggers, NewConsoleLogger(owner))
	}

	if enableFile {
		name := ""
		if fileName != nil {
			name = *fileName
		} else {
			name = owner + time.Now().String() + ".log"
		}
		loggers = append(loggers, newFileLogger(name, owner))
	}

	if enableSocket {
		sLogger, err := newSocketLogger(owner, port, enableSocket)
		if err != nil {
			return nil, err
		}

		loggers = append(loggers, sLogger)
	}
	return &CompositeLogger{loggers: loggers}, nil
}

func (c *CompositeLogger) Info(message string) {
	for _, l := range c.loggers {
		l.Info(message)
	}
}

func (c *CompositeLogger) Debug(message string) {
	for _, l := range c.loggers {
		l.Debug(message)
	}
}

func (c *CompositeLogger) Warning(message string) {
	for _, l := range c.loggers {
		l.Warning(message)
	}
}

func (c *CompositeLogger) Error(message string) {
	for _, l := range c.loggers {
		l.Error(message)
	}
}

func (c *CompositeLogger) logInner(logFor func(Logger)) {
	for _, l := range c.loggers {
		logFor(l)
	}
}

func (c *CompositeLogger) SetWriteState(state bool) {
	for _, l := range c.loggers {
		l.SetWriteState(state)
	}
}

func (c *CompositeLogger) Dispose() {
	for _, l := range c.loggers {
		l.Dispose()
	}
}
