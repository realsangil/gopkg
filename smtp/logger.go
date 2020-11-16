package lalasmtp

// Logger is the fundamental interface for logging.
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

func print(logger Logger, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Print(args...)
}

func printf(logger Logger, format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Printf(format, args...)
}
