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
	logFatal uint8 = 1 << iota
	logError
	logWarning
	logInfo
	logDebug
	logCustom
)

var (
	logBitmask    uint8
	defaultWriter io.Writer
)

// Initializes the package with default settings
func init() {

	defaultWriter = os.Stderr
	//logBitmask = logFatal | logError | logWarning | logInfo | logDebug | logCustom
	TurnOnAllLogging()
}

// This is the only logging that wont return a bool if it worked as it will
// work similarly to log.Fatal
func Fatal(format string, args ...interface{}) {
	formatter(defaultWriter, logFatal, "FATAL", format, args...)
	os.Exit(1)
}

func Error(format string, args ...interface{}) bool {
	return formatter(defaultWriter, logError, "ERROR", format, args...)
}

func Warn(format string, args ...interface{}) bool {
	return formatter(defaultWriter, logWarning, "WARN", format, args...)
}

func Info(format string, args ...interface{}) bool {
	return formatter(defaultWriter, logInfo, "INFO", format, args...)
}

func Debug(format string, args ...interface{}) bool {
	return formatter(defaultWriter, logDebug, "DEBUG", format, args...)
}

// A custom logger that can be used
func Custom(writer io.Writer, logLevelString string, format string, args ...interface{}) bool {
	// The numerical log level for a custom log message is 1
	return formatter(writer, logCustom, logLevelString, format, args...)
}

// Returns true on success and false on block
func formatter(writer io.Writer, numericalLogLevel uint8, logLevelString string, format string, args ...interface{}) bool {

	if numericalLogLevel&logBitmask == 0 {
		return false
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	fileName, caller, line := getDetails()

	message := fmt.Sprintf(format, args...)

	// <date and time> [<log level>] <filePath>:<line number>:<caller>() <formatted message>\n
	fmt.Fprintf(writer, "%s [%s] %s:%d:%s() %s\n", now, logLevelString, fileName, line, caller, message)
	return true
}

// This function retrieves the file which called the
// logging funtion and which line from where it was called
func getDetails() (string, string, int) {
	pc, path, line, ok := runtime.Caller(3)

	// Something went wrong
	if !ok {
		return "???", "???", -1
	}

	paths := strings.Split(path, "/")
	file := paths[len(paths)-1]

	caller := runtime.FuncForPC(pc).Name()
	stack := strings.Split(caller, ".")
	caller = stack[len(stack)-1]

	return file, caller, line
}

// Allows for changing of the default io.Writer that the logger uses
func SetDefaultWriter(newWriter io.Writer) {
	defaultWriter = newWriter
}

// A log gets printed if the bitmask 'allows it'.
// This function allows the user to quickly update the entire bitmask
func SetLogBitmask(bitmask uint8) uint8 {
	logBitmask = bitmask
	return logBitmask
}

func TurnOnAllLogging() uint8 {
	logBitmask = logFatal | logError | logWarning | logInfo | logDebug | logCustom
	return logBitmask
}

// These functions allows the user to toggle each type of log individually
// Returns the logBitmask after the chang
func SetLogFatal(b bool) uint8 {
	logToggle(b, logFatal)
	return logBitmask
}
func SetLogError(b bool) uint8 {
	logToggle(b, logError)
	return logBitmask
}
func SetLogWarning(b bool) uint8 {
	logToggle(b, logWarning)
	return logBitmask
}
func SetLogInfo(b bool) uint8 {
	logToggle(b, logInfo)
	return logBitmask
}
func SetLogDebug(b bool) uint8 {
	logToggle(b, logDebug)
	return logBitmask
}
func SetLogCustom(b bool) uint8 {
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
