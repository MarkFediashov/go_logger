package go_logger

const format = "|%-23s| %-6s [%-15s]: %v\n"
const sourceFieldLength = 15

type Logger interface {
	Debug(string)
	Warning(string)
	Error(string)
	Info(string)

	SetWriteState(bool)
	Dispose()
}

type Config struct {
	EnableConsole bool
	EnableFile    bool
	EnableSocket  bool
}

func NewLogger(withConsole bool, fileLoggingName *string, outPort *int) Logger {
	return nil
}
