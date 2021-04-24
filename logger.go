package bord

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// Here are the various types of logs that can be used
const (
	logCustom uint8 = 1 << iota
	logError
	logWarning
	logInfo
	logDebug
)

var (
	logBitmask    uint8
	defaultWriter io.Writer
)

// Initializes the package with default settings
func init() {

	defaultWriter = os.Stderr
	logBitmask = logCustom | logError | logWarning | logInfo | logDebug
}

func Error(format string, args ...interface{}) {
	formatter(defaultWriter, logError, "ERROR", format, args...)
}

func Warn(format string, args ...interface{}) {
	formatter(defaultWriter, logWarning, "WARN", format, args...)
}

func Info(format string, args ...interface{}) {
	formatter(defaultWriter, logInfo, "INFO", format, args...)
}

func Debug(format string, args ...interface{}) {
	formatter(defaultWriter, logDebug, "DEBUG", format, args...)
}

// A custom logger that can be used
func Custom(writer io.Writer, logLevelString string, format string, args ...interface{}) {
	// The numerical log level for a custom log message is 1
	formatter(writer, 1, logLevelString, format, args...)
}

func formatter(writer io.Writer, numericalLogLevel uint8, logLevelString string, format string, args ...interface{}) {

	if numericalLogLevel&logBitmask == 0 {
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	fileName, line := getDetails()

	message := fmt.Sprintf(format, args...)

	// <date and time> [<log level>] .\<filePath>:<line number> <formatted message>\n
	fmt.Fprintf(writer, "%s [%s] .\\%s:%d %s\n", now, logLevelString, fileName, line, message)
}

// This function retrieves the file which called the
// logging funtion and which line from where it was called
func getDetails() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	files := strings.Split(file, "/")
	file = files[len(files)-1]

	return file, line
}

// Allows for changing of the default io.Writer that the logger uses
func SetDefaultWriter(newWriter io.Writer) {
	defaultWriter = newWriter
}

// A log gets printed if the bitmask 'allows it'.
// This function allows the user to quickly update the entire bitmask
func SetLogBitmask(bitmask uint8) {
	logBitmask = bitmask
}

// These functions allows the user to toggle each type of log individually
// Returns the logBitmask after the chang
func LogError(b bool) uint8 {
	logToggle(b, logError)
	return logBitmask
}
func LogWarning(b bool) uint8 {
	logToggle(b, logWarning)
	return logBitmask
}
func LogInfo(b bool) uint8 {
	logToggle(b, logInfo)
	return logBitmask
}
func LogDebug(b bool) uint8 {
	logToggle(b, logDebug)
	return logBitmask
}
func LogCustom(b bool) uint8 {
	logToggle(b, logCustom)
	return logBitmask
}

func logToggle(b bool, logType uint8) {
	if b {
		logBitmask |= logType
	} else {
		// Clear the logType bit from the LogBitmask
		logBitmask &^= logType
	}
}
