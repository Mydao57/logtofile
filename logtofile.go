// This file contains the implementation of a logging package that writes log messages to a file.
// The package provides functions to log messages with different severity levels such as INFO, ERROR, DEBUG, and FATAL.
// The log messages are written to a file with a timestamp and the severity level.
// Author: Mydao57

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

// log writes a log message to the file with the specified severity and message.
func (l *Logger) log(severity, message string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [%s] %s\n", now, severity, message)
	l.file.WriteString(logMessage)
}

// INFO logs an information message to the log file.
func INFO(message string) {
	if logInstance != nil {
		logInstance.log("INFO", message)
	}
}

// ERROR logs an error message to the log file.
func ERROR(message string) {
	if logInstance != nil {
		logInstance.log("ERROR", message)
	}
}

// DEBUG logs a debug message to the log file.
func DEBUG(message string) {
	if logInstance != nil {
		logInstance.log("DEBUG", message)
	}
}

// FATAL logs a fatal error message to the log file.
func FATAL(message string) {
	if logInstance != nil {
		logInstance.log("FATAL", message)
	}
}

// Close the log file
func Close() error {
	if logInstance != nil {
		return logInstance.file.Close()
	}
	return nil
}