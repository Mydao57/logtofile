package logtofile

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var logInstance *Logger

type Logger struct {
	file *os.File
}

func NewLogger(directory string) error {
	// Ensure the directory exists
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create log directory: %v", err)
	}

	now := time.Now()

	// Generate the log file name
	fileName := fmt.Sprintf("log_%04d%02d%02d_%02d%02d%02d.log", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	filePath := filepath.Join(directory, fileName)

	// Create the log file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not create log file: %v", err)
	}

	logInstance = &Logger{file: file}
	return nil
}
