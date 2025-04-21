// logging class, created upon init.

package internal

import (
	"log"
	"os"
)

var (
	Action *log.Logger
	Warn   *log.Logger
	Info   *log.Logger
	Error  *log.Logger
	Debug  *log.Logger
)

var reset = "\033[0m"
var red = "\033[31m"
var yellow = "\033[33m"
var magenta = "\033[35m"

func init() {
	Info = log.New(os.Stdout, "", 0)
	Warn = log.New(os.Stdout, yellow+"warn: "+reset, 0)
	Error = log.New(os.Stdout, red+"error: "+reset, 0)
	Debug = log.New(os.Stdout, magenta+"debug: "+reset, 0)
}
