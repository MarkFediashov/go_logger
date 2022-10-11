package go_logger

import (
	"os"
	"sync"
	"syscall"
)

type fileLogger struct {
	loggerState
	fileName   string
	file       *os.File
	writeState sync.Mutex
}

func newFileLogger(fileName string, owner string) *fileLogger {
	file, err := os.OpenFile(fileName, os.O_CREATE|syscall.O_APPEND, 0777)
	if err != nil {
		return nil
	}
	return &fileLogger{
		fileName: fileName,
		file:     file,
		loggerState: loggerState{
			state: true,
			owner: owner,
		},
	}
}

func (logger *fileLogger) Info(message string) {
	logger.logInner(INFO, message)
}

func (logger *fileLogger) Debug(message string) {
	logger.logInner(DEBUG, message)
}

func (logger *fileLogger) Warning(message string) {
	logger.logInner(WARNING, message)
}

func (logger *fileLogger) Error(message string) {
	logger.logInner(ERROR, message)
}

func (logger *fileLogger) logInner(level logLevel, message string) {
	if logger.state {
		record := formatLogString(level, logger.owner, message)
		go logger.writeGoroutine(record)
	}
}

func (logger *fileLogger) writeGoroutine(row string) {
	logger.writeState.Lock()
	defer logger.writeState.Unlock()

	logger.file.WriteString(row)
}

func (logger *fileLogger) SetWriteState(state bool) {
	logger.state = state
}

func (logger *fileLogger) Dispose() {
	logger.file.Close()
}
