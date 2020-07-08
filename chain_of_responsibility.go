package main

import "fmt"

const (
	INFO  = 1
	DEBUG = 2
	ERROR = 3
)

type ConsoleLogger struct {
	Level      int
	NextLogger interface{}
}

func NewConsoleLogger(level int) *ConsoleLogger {
	return &ConsoleLogger{
		Level: level,
	}
}

func (a *ConsoleLogger) LogMessage(level int, message string) {
	if a.Level == level {
		a.Write(message)
		return
	}
	if a.NextLogger != nil {
		switch v := a.NextLogger.(type) {
		case *ErrorLogger:
			v.LogMessage(level, message)
		case *ConsoleLogger:
			v.LogMessage(level, message)
		case *FileLogger:
			v.LogMessage(level, message)
		}
	}
}

func (c *ConsoleLogger) Write(message string) {
	fmt.Println("stdout: ", message)
}

type ErrorLogger struct {
	Level      int
	NextLogger interface{}
}

func NewErrorLogger(level int) *ErrorLogger {
	return &ErrorLogger{
		Level: level,
	}
}

func (a *ErrorLogger) LogMessage(level int, message string) {
	if a.Level == level {
		a.Write(message)
		return
	}
	if a.NextLogger != nil {
		switch v := a.NextLogger.(type) {
		case *ErrorLogger:
			v.LogMessage(level, message)
		case *ConsoleLogger:
			v.LogMessage(level, message)
		case *FileLogger:
			v.LogMessage(level, message)
		}
	}
}

func (e *ErrorLogger) Write(message string) {
	fmt.Println("stderr: ", message)
}

type FileLogger struct {
	Level      int
	NextLogger interface{}
}

func NewFileLogger(level int) *FileLogger {
	return &FileLogger{
		Level: level,
	}
}

func (a *FileLogger) LogMessage(level int, message string) {
	if a.Level == level {
		a.Write(message)
		return
	}
	if a.NextLogger != nil {
		switch v := a.NextLogger.(type) {
		case *ErrorLogger:
			v.LogMessage(level, message)
		case *ConsoleLogger:
			v.LogMessage(level, message)
		case *FileLogger:
			v.LogMessage(level, message)
		}
	}
}

func (f *FileLogger) Write(message string) {
	fmt.Println("file logger: ", message)
}

func main() {
	console := &ConsoleLogger{
		Level: INFO,
	}
	err := &ErrorLogger{
		Level: ERROR,
	}
	file := &FileLogger{
		Level: DEBUG,
	}

	chain := &ConsoleLogger{
		Level:      INFO,
		NextLogger: console,
	}
	chain.NextLogger.(*ConsoleLogger).NextLogger = err
	chain.NextLogger.(*ConsoleLogger).NextLogger.(*ErrorLogger).NextLogger = file

	chain.LogMessage(DEBUG, "This is a debug info")
	chain.LogMessage(ERROR, "This is an error info")
	chain.LogMessage(INFO, "This is an info")
}
