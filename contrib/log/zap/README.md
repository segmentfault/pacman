# log
> The logging implemented by[zap](http://go.uber.org/zap)

## Usage
```go
import (
    "github.com/segmentfault/pacman/contrib/log/zap"
    "github.com/segmentfault/pacman/log"
)
```

```go
    log.SetLogger(zap.NewLogger(log.ParseLevel(logLevel), zap.WithName(Name), zap.WithPath(logPath), zap.WithCallerFullPath()))
    log.Debug("test")
    log.Info("test")
    log.Warn("test")
    log.Error("test")
```
