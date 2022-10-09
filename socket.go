package logger

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
	result := &socketLogger{loggerState: loggerState{owner: owner, state: enableSocket}, listener: listener, port: port}

	go result.acceptByChannel()
	go result.acceptConnectionsBackground()

	return result, nil
}

func (s *socketLogger) acceptByChannel() {
	for {
		connection, err := s.listener.Accept()
		if err != nil {
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

func (s *socketLogger) Log(data ...any) {
	if s.state {
		row := formatLogString(s.owner, data)
		bytes := []byte(row)
		for _, s := range s.clients {
			go s.Write(bytes)
		}
	}
}

func (s *socketLogger) SetWriteState(state bool) {
	s.state = state
}

func (s *socketLogger) Dispose() {
	s.listener.Close()
}
