package main

import (
	"errors"
	"fmt"
)

type FileLogger struct{}

func (fl FileLogger) Info(msg string) error {
	fmt.Println("FileLogger Info:", msg)
	return nil
}

func (fl FileLogger) Warning(msg string) error {
	fmt.Println("FileLogger Warning:", msg)
	return nil
}

func (fl FileLogger) Error(msg string) error {
	fmt.Println("FileLogger Error:", msg)
	return nil
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Info(msg string) error {
	fmt.Println("ConsoleLogger Info:", msg)
	return nil
}

func (cl ConsoleLogger) Warning(msg string) error {
	fmt.Println("ConsoleLogger Warning:", msg)
	return nil
}

func LogMessage(loggers []Logger, msg string) error {
	if len(loggers) == 0 {
		return errors.New("no loggers provided")
	}

	for _, logger := range loggers {
		_ = logger.Info(msg)
		_ = logger.Warning(msg)
		_ = logger.Error(msg)
	}

	return nil
}

func main() {
	fl := FileLogger{}
	cl := ConsoleLogger{}

	loggers := []Logger{fl, cl}

	_ = LogMessage(loggers, "Hello world!")
}
