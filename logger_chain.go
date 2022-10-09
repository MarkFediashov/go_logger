package logger

import (
	"time"
)

type ChainedLogger struct {
	loggers []Logger
	console *consoleLogger
	file    *fileLogger
	network *socketLogger
}

func NewChainedLogger(owner string, fileName *string, enableConsole bool, enableFile bool, enableSocket bool, port int) (*ChainedLogger, error) {

	loggers := make([]Logger, 0, 3)

	if enableConsole {
		loggers = append(loggers, &consoleLogger{loggerState: loggerState{owner: owner, state: enableConsole}})
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
	return &ChainedLogger{loggers: loggers}, nil
}

func (c *ChainedLogger) Log(message string) {
	for _, l := range c.loggers {
		l.Log(message)
	}
}

func (c *ChainedLogger) SetWriteState(consoleState bool, fileState bool, socketState bool) {
	c.console.state = consoleState
	c.file.state = fileState
	c.network.state = socketState
}

func (c *ChainedLogger) Dispose() {
	for _, l := range c.loggers {
		l.Dispose()
	}
}
