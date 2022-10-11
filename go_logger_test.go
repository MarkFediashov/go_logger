package go_logger

import (
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	filename := "test.log"
	port := 47891
	logger, err := NewChainedLogger("Test", &filename, true, true, true, port)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	} else {
		fmt.Println("Logger created")
	}

	connection, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	buffer := make([]byte, 1024)
	logger.Warning("Test)")
	logger.Error("Test)")
	logger.Debug("Test)")
	logger.Info("Test)")
	_, err = connection.Read(buffer)
	if err == nil {
		fmt.Println(string(buffer))
	}
	time.Sleep(2)
}
