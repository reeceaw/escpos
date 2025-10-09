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
	// Create io.Writer where commands will be written to
	file, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create Client with profile for your target printer
	client := escpos.NewClient(file, escpos.EpsonTMT20III{})
	
	// Send text to print!
	client.WriteLine("My first line")
	client.WriteLine("My second line")
	client.WriteLine("My third line")
	
	// Cut paper and finish printing
	client.Cut()
	client.End()
}
```

### Profiles
Printer-agnostic functions are provided by the `Client`, such as `Cut()` - all ESC/POS printers
should have a cut function. A `Profile` is used to map from these agnostic functions to the
printer-specific commands.

**The only profile provided out-of-the-box is for the Epson TM-T20III, shown in the example above.**

Should you wish to use this library with a different printer but you find the TM-T20III profile
to be incompatible, you can implement the `Profile` interface (rather all the interfaces it groups)
with the commands that are specific to your printer. The [Epson ESC/POS Command Reference](https://download4.epson.biz/sec_pubs/pos/reference_en/escpos/index.html)
specifies commands and their support in different models.

### Formatting
You can apply formatting by using the `Write(string, FormatConfig)` function:
```go
func writeWithFormatting(client escpos.Client) {
	customFormat := escpos.DefaultFormatConfig().
		Emphasize(true).
		Font("B").
		Justify("center").
		Underline("1-dot")
	
	client.Write("My formatted message!\n", customFormat)
}
```
