# Bord

Bord is a logger for the [GO](https://golang.org/) programming language

## Usage

Using Bord is simple as its possible to import the package anyhwere it is required 

Example syntax is:
```go
bord.Error("This is an error message: '%s'", err)
bord.Warn("This is a warning message '%s'", "<warning message>")
bord.Info("This is an info message '%s'", "<info message>")
bord.Debug("This is a debug message '%s'", "<debug message>")
bord.Custom(os.Stderr, "CUSTOM", "This is a %s message with a custom log level tag", "custom")
```

## Options

To customize what logging messages that get displayd can the following syntax be used
```go
// To turn on indivudual logging
LogError(true)
LogWarning(true)
LogInfo(true)
LogDebug(true)
LogCustom(true)

// To turn off indivudual logging
LogError(false)
LogWarning(false)
LogInfo(false)
LogDebug(false)
LogCustom(false)
```

A bitmask is used to calcualte what get showed so
custom = 1
error = 2
warning = 4
info = 8
debug = 16

This allows you to input a prepared value to toggle all logging that is desired
```go
bord.SetLogBitmask(31) // Will turn on everything
bord.SetLogBitmask(27) // Will turn on everything except warnings
bord.SetLogBitmask(0) // Turns off all logging
```

## Test coverage

Todo