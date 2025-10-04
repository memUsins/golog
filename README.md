# GoLog

golog — adapter based logger util

## Installation

```shell
$ go get github.com/memUsins/golog
```

## Adapters

| Adapter                                     | destination         |
|---------------------------------------------|---------------------|
| [Console](github.com/memUsins/gologconsole) | Console adapter     |
| [File](github.com/memUsins/gologfile)       | File (json) adapter |

## Usage

### Logging Methods with Suffixes
- **Methods with the `n` suffix** allow specifying an additional `name` field: `logger.Infon("TestLogger", "Some log info")`
- **Methods with the `f` suffix** allow formatted string logging: `logger.Infof("%s string", "Some formatted")`

### Additional Methods for Logging
- **DebugJSON** for formatting JSON objects: `logger.DebugJSON(myStruct)`
- **WithFields**, **WithError**, **WithName** for detailed context:
  ```go
  logger.WithFields(golog.F{"foo": "bar"}).Info("Some log info") // golog.F = map[string]interface{}
  logger.WithError(errors.New("my error")).Info("Some log info")
  logger.WithName("TestLogger").Info("Some log info")
  ```
