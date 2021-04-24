# Bord

Bord is a logger for the [GO](https://golang.org/) programming language

100% test coverage

## Install

```
go get github.com/CarlFlo/bord
```

The main functionality is working, but i'm always working to improve the package and fix bugs.

## Usage

The goal of Board is to make powerful logging simple.

You are able to get started with no configuration required.

By default will Bord output to **os.Stderr** and output all types of loggin


Example syntax is:
```go
bord.Error("This is an error message: '%s'", err)
bord.Warn("This is a warning message '%s'", "<warning message>")
bord.Info("This is an info message '%s'", "<info message>")
bord.Debug("This is a debug message '%s'", "<debug message>")
bord.Custom(os.Stderr, "CUSTOM", "This is a %s message with a custom log level tag", "custom")
```

Each of the functions will return **true** if the message could be logged, i.e. printed to designated io.Writer, and **false** if it was blocked.

## Options & Customization

To customize what logging messages that get displayed can the following syntax be used to change the settings

Each of these functions will return the updated bitmap (uint8)

```go
// To turn on indivudual logging
bord.SetLogError(true)
bord.SetLogWarning(true)
bord.SetLogInfo(true)
bord.SetLogDebug(true)
bord.SetLogCustom(true)

// To turn off indivudual logging
bord.SetLogError(false)
bord.SetLogWarning(false)
bord.SetLogInfo(false)
bord.SetLogDebug(false)
bord.SetLogCustom(false)
```

A bitmask is used to calcualte what get showed:
* custom = 1
* error = 2
* warning = 4
* info = 8
* debug = 16

This allows you to input a prepared value to toggle the desired logging.
```go
bord.SetLogBitmask(31) // Will turn on everything
bord.SetLogBitmask(27) // Will turn on everything except warnings (4)
bord.SetLogBitmask(0) // Turns off all logging
```

The default output is **os.Stderr**, but this can be changed with
```go
bord.SetDefaultWriter(newWriter) // Will accept any io.Writer
```