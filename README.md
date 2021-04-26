# Bord

Bord is a simple to use logger for the [GO](https://golang.org/) programming language

Test coverage: **93.3%**

## Features
- Detailed logging showing the message origin file, calling function and line.
- Customizable. Select what types of logging messages you want
- Simplistic. No configuration to get started and easy to change how you want it


## Install

```
  go get github.com/CarlFlo/bord
```

The main functionality is working, but i'm always working to improve the package and fix bugs.

## Usage

The goal of Board is to make powerful logging simple.

You are able to get started with no configuration required.

By default will Bord output **all types** of log messages to **os.Stderr**.


Example syntax is:
```go
bord.Fatal("This is an fatal log message: '%s'", err) // Will os.Exit(1)
bord.Error("This is an error message: '%s'", err)
bord.Warn("This is a warning message '%s'", "<warning message>")
bord.Info("This is an info message '%s'", "<info message>")
bord.Debug("This is a debug message '%s'", "<debug message>")
bord.Custom(os.Stderr, "CUSTOM", "This is a %s message with a custom log tag", "custom")
bord.Custom(os.Stderr, "NETWORK", "Another example with a different log tag")
```

Each of the logging functions above will return **True** if the *message could be logged* and **False** if it was blocked by a setting.

## Options & Customization

To customize the logging messages that get displayed can the following syntax be used to change the settings.

```go
// To turn on indivudual logging
bord.SetLogFatal(true)
bord.SetLogError(true)
bord.SetLogWarning(true)
bord.SetLogInfo(true)
bord.SetLogDebug(true)
bord.SetLogCustom(true)

// To turn off indivudual logging
bord.SetLogFatal(false)
bord.SetLogError(false)
bord.SetLogWarning(false)
bord.SetLogInfo(false)
bord.SetLogDebug(false)
bord.SetLogCustom(false)
```
Each of the above functions will return the updated bitmap (uint8)

A bitmask is used to calcualte what get outputted:
* fatal = 1
* error = 2
* warning = 4
* info = 8
* debug = 16
* custom = 32

This allows the user to input a prepared value, i.e. from a configuration file, to set the desired logging.
```go
bord.SetLogBitmask(63) // Will turn on everything 1+2+4+8+16+32=63
bord.SetLogBitmask(59) // Will turn on everything except warnings (4) 1+2+8+16+32=59
bord.SetLogBitmask(0) // Turns off all logging
```

The default output is **os.Stderr**, but this can be changed with:
```go
bord.SetDefaultWriter(newWriter) // Will accept any io.Writer
```

## TODO

- [ ] Option for 'verbosity'. Being able to change verbose level for different type of log message types
