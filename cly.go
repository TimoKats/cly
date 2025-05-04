// main control flow of cly. Gets CLI args and calls functions

package main

import (
	"errors"

	cly "github.com/TimoKats/cly/internal"

	"os"
)

func run(config cly.Config, args []string) error {
	if len(args) < 3 {
		return errors.New("no command to run provided")
	}
	config.AddArgs(args)
	alias, aliasFound := config.GetAlias(args, 2)
	if aliasFound {
		if alias.Concurrent {
			return alias.ConcurrentRun()
		}
		return alias.SequentialRun()
	}
	return errors.New("alias not found in yaml")
}

func parse(config cly.Config, args []string) error {
	switch {
	case len(args) == 1:
		return errors.New("no command found")
	case args[1] == "run":
		return run(config, args)
	case args[1] == "ls":
		return config.List(args)
	default:
		return errors.New("command '" + args[1] + "' not valid")
	}
}

func main() {
	var config cly.Config
	var err error
	if config, err = cly.Parse(); err != nil {
		cly.Error.Println(err)
		return
	}
	if err = parse(config, os.Args); err != nil {
		cly.Error.Println(err)
	}
}
