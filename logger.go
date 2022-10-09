package logger

const format = "|%s| [%-15s]: %v\n"
const sourceFieldLength = 15

type Logger interface {
	Log(string)
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
