# escpos
A Go library for interacting with ESC/POS printers.
> ⚠️ This library is under development.

### Getting started
Add `escpos` to your module:
```shell
go get github.com/reeceaw/escpos
```
Import `escpos` and use `NewClient(io.Writer)` to create a client. 
```go
package main

import (
	"os"
	"github.com/reeceaw/escpos"
)

func main() {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	client := escpos.NewClient(f)
	client.WriteLine("My first line")
	client.WriteLine("My second line")
	client.WriteLine("My third line")
	client.Cut()
	client.End()
}
```