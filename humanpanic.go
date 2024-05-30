package humanpanic

import (
	"fmt"
	"os"
	"runtime"
)

// HumanError represents an error caused by a human.
type HumanError struct {
	msg   string
	file  string
	line  int
	trace []byte
}

// Error implements the error interface.
func (e *HumanError) Error() string {
	return e.msg
}

func (e *HumanError) String() string {
	return fmt.Sprintf(
		"Recovered from error: %s\nFile: %s:%d\nStack trace:\n%s",
		e.msg, e.file, e.line, e.trace)
}

func (e *HumanError) logToFile() error {
	f, err := os.CreateTemp("", "human-panic-*.log")
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(e.String()); err != nil {
		return err
	}

	fmt.Printf("Error details written to file: %s\n", f.Name())
	return nil
}

func Panic(msg string) {
	_, file, line, _ := runtime.Caller(1)
	trace := make([]byte, 8192)
	count := runtime.Stack(trace, false)
	trace = trace[:count]

	err := &HumanError{
		msg:   msg,
		file:  file,
		line:  line,
		trace: trace,
	}

	panic(err)
}

func Panicf(format string, args ...interface{}) {
	Panic(fmt.Sprintf(format, args...))
}

// Recover regains control of a panicking goroutine.
// It returns the error that caused the panic, or nil if no error occurred.
func Recover() error {
	if err := recover(); err != nil {
		humanErr, ok := err.(*HumanError)
		if !ok {
			return humanErr
		}
		fmt.Println("Well, That's embarrassing. To help us diagnose the problem you can send us a crash report")
		fmt.Print(humanErr.String())
		if logErr := humanErr.logToFile(); logErr != nil {
			return fmt.Errorf("failed to log error to file: %w", logErr)
		}
		return humanErr
	}
	return nil
}
