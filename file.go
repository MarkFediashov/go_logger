package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type fileLogger struct {
	loggerState
	fileName   string
	file       *os.File
	writeState sync.Mutex
}

func newFileLogger(fileName string, owner string) *fileLogger {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, os.ModeAppend)
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

func (logger *fileLogger) Log(data ...any) {
	if logger.state {
		record := fmt.Sprintf(format, time.Now().Format("2006-01-02T15:04:05.999"), logger.owner, data)
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
