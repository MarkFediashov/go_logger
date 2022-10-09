package logger

import "net"

type LoggerOutput struct {
	isAttached bool
	conn       net.Conn
}

func NewLoggerOutput() *LoggerOutput {
	return &LoggerOutput{isAttached: false}
}

func (l *LoggerOutput) Attach(addr string) {

}
