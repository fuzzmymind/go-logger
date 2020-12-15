# go-logger

A go wrapper around the awesome zap and lumberjack library to create json log files for your application.

### Usage

##### In your main package (entrypoint)

```
package main

import (
	"github.com/fuzzmymind/go-logger/logger"
)

func main() {
	lc := logger.Config{
		FileName:   "test",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     1,
		LogLevel:   logger.InfoLevel,
	}
	logger.Init(lc)
	logger.Info("Hello", logger.Field("Key", 123))
}
```

##### In your other packages of your application, no need to use init. Just import and log.

```
package somepackage

import (
	"github.com/fuzzmymind/go-logger/logger"
)

func somemethod() {
	logger.Info("Hello", logger.Field("Key", 123))
}
```
