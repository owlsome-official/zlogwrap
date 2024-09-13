# zlogwrap

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

zlogwrap is a logging library, which wrapping all common use from zerolog.

## Table of Contents

- [zlogwrap](#zlogwrap)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Signatures](#signatures)
    - [zerologWrapper](#zerologwrapper)
  - [Examples](#examples)
  - [Config](#config)
  - [Default Config](#default-config)
  - [Dependencies](#dependencies)
  - [More Example Usage](#more-example-usage)

## Installation

```bash
  go get -u github.com/owlsome-official/zlogwrap
```

## Signatures

```go
func New(conf ...Config) zerologWrapper
```

### zerologWrapper

Interface of functions

```go
type zerologWrapper interface {
  SetField(key string, anything interface{}) zerologWrapper // Set field in logs
  Debug(anything ...interface{})                            // level 0
  Info(anything ...interface{})                             // level 1
  Warn(anything ...interface{})                             // level 2
  Error(anything ...interface{})                            // level 3
  Fatal(anything ...interface{})                            // level 4
  Panic(anything ...interface{})                            // level 5
  GetLogEvent(zerolog.Level) *zerolog.Event                 // With Caller (file:line)
}
```

## Examples

Import the package

```go
import (
  "github.com/owlsome-official/zlogwrap"
)
```

Using with these

```go
logger := zlogwrap.New()
logger.Debug("A Debug Log")

// {"level":"debug","time":"2021-07-27T16:49:59+07:00","message":"A Debug Log"}
```

Optional: With configuration

```go
logger := zlogwrap.New(zlogwrap.Config{
  ServiceName: "zlogwrap",
})
logger.Debug("A Debug Log")
logger.SetField("key", "value").Debug("A Debug Log")

// {"level":"debug","service":"zlogwrap","time":"2021-08-05T11:14:00+07:00","message":"A Debug Log"}
// {"level":"debug","key":"value","service":"zlogwrap","time":"2021-08-05T11:14:00+07:00","message":"A Debug Log"}
```

or customize the "Logger" and this library support middleware [RequestID](https://github.com/gofiber/docs/blob/master/api/middleware/requestid.md)

```go
myLogger := zerolog.New(os.Stdout).
  With().
  Str("foo", "bar").
  Float64("money", 10.99).
  Logger()
logger := zlogwrap.New(zlogwrap.Config{
  ServiceName: serviceName,
  Logger:      myLogger,
  Context:     c, // Which type is `*fiber.Ctx`
})
logger.Debug("A Debug Log")

// ps. see the real usage in `example/main.go`
// {"level":"debug","foo":"bar","money":10.99,"service":"zlogwrap","transaction_id":"b91f468c-c608-4729-88d8-9a12c951c31a","url":"/","message":"A Debug Log"}
```

If you need terminate app when logging fatal / panic level, add a config flag.

```go
logger := zlogwrap.New(zlogwrap.Config{
  TerminateOnCriticalLevel: true,
})
logger.Panic("foo")
```

## Config

```go
// Config defines the config for library.
type Config struct {

  // Optional. Default: false
  Hidden bool

  // Optional. Default: ""
  ServiceName string

  // Optional. Default: false
	TerminateOnCriticalLevel bool

  // Optional. Default: &fiber.Ctx{}
  Context *fiber.Ctx

  // Optional. Default: log.Logger
  Logger zerolog.Logger
}
```

## Default Config

```go
var ConfigDefault = Config{
	Hidden:                   false,
	ServiceName:              "",
	TerminateOnCriticalLevel: false,
	Context:                  nil,
	Logger:                   log.Logger,
}
```

## Dependencies

- [Zerolog](https://github.com/rs/zerolog)
- [Fiber](https://github.com/gofiber/fiber)

## More Example Usage

Please go to [example/main.go](./example/main.go)
