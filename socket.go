package logger

import "net"

type socketLogger struct {
	loggerState
	port     int
	listener []net.Listener
}

func (s *socketLogger) newSocketLogger() {

}

func (s *socketLogger) Log(data ...any) {

}

func (s *socketLogger) SetWriteState(state bool) {

}

func (s *socketLogger) Dispose() {

}
