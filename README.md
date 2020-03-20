# go-logif

Golang: The abstract interface of a log output, and the wrapper of a standard "log" package.

1) Golang standard (Print/Fatal/Panic)
2) "Print" with message level (Debug/Info/Warn/Error)

## Usage

```golang
package main

import "github.com/shimt/go-logif/gologif"

func main() {
    l := gologif.New(os.Stderr, "", gologif.LstdFlags)
	l.SetOutputLevel(gologif.DEBUG)
	l.Print("message")
	l.Debug("debug message")
	l.Info("info message")
	l.Warn("warn message")
	l.Error("error message")
    // Output:
    // 2020/03/22 14:06:21 message
    // 2020/03/22 14:06:21 [DEBUG] debug message
    // 2020/03/22 14:06:21 [INFO] info message
    // 2020/03/22 14:06:21 [WARN] warn message
    // 2020/03/22 14:06:21 [ERROR] error message
}
```

