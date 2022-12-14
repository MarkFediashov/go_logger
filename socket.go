package go_logger

import (
	"net"
	"strconv"
)

type socketLogger struct {
	loggerState
	listener net.Listener
	port     int
	clients  []net.Conn
	dispose  chan struct{}
	connChan chan net.Conn
}

func newSocketLogger(owner string, port int, enableSocket bool) (*socketLogger, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	result := &socketLogger{loggerState: loggerState{owner: owner, state: enableSocket}, listener: listener, port: port,
		dispose:  make(chan struct{}),
		connChan: make(chan net.Conn),
	}

	go result.acceptByChannel()
	go result.acceptConnectionsBackground()

	return result, nil
}

func (s *socketLogger) acceptByChannel() {
	for {
		connection, err := s.listener.Accept()
		if err == nil {
			s.connChan <- connection
		} else {
			close(s.connChan)
		}
	}

}

func (s *socketLogger) acceptConnectionsBackground() {
	var c net.Conn
	var ok bool
	for {
		c, ok = <-s.connChan
		if ok {
			println("Connection accepted!")
			s.clients = append(s.clients, c)
		} else {
			println("Channel was closed")
			return
		}

	}
}

func (s *socketLogger) Info(message string) {
	s.logInner(INFO, message)
}

func (s *socketLogger) Debug(message string) {
	s.logInner(DEBUG, message)
}

func (s *socketLogger) Warning(message string) {
	s.logInner(WARNING, message)
}

func (s *socketLogger) Error(message string) {
	s.logInner(ERROR, message)
}

func (s *socketLogger) logInner(level logLevel, message string) {
	if s.state {
		row := formatLogString(level, s.owner, message)
		bytes := []byte(row)
		for _, c := range s.clients {
			go c.Write(bytes)
		}
	}
}

func (s *socketLogger) SetWriteState(state bool) {
	s.state = state
}

func (s *socketLogger) Dispose() {
	s.listener.Close()
}
