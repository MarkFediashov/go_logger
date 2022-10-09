package logger

type ChainedLogger struct {
	loggers []Logger
}

func (c *ChainedLogger) Log(data ...any) {
	for _, l := range c.loggers {
		l.Log(data)
	}
}

func (c *ChainedLogger) SetWriteState(console bool, file bool, socket bool) {
	for _, l := range c.loggers {
		switch l.(type) {
		case *consoleLogger:
			l.SetWriteState(console)
		case *fileLogger:
			l.SetWriteState(file)
		case *socketLogger:
			l.SetWriteState(socket)
		}
	}
}

func (c *ChainedLogger) Dispose() {
	for _, l := range c.loggers {
		l.Dispose()
	}
}
