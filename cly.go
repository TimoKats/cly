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
		return alias.Run()
	}
	return errors.New("alias not found in yaml")
}

func ls(config cly.Config, arg string) error {
	if arg == "ls" {
		return config.List(false)
	}
	return config.List(true)
}

func parse(config cly.Config, args []string) error {
	switch {
	case len(args) == 1:
		return errors.New("no command found")
	case args[1] == "run":
		return run(config, args)
	case args[1] == "ls", args[1] == "tree":
		return ls(config, args[1])
	default:
		return errors.New("command '" + args[1] + "' not valid")
	}
}

func main() {
	config, err := cly.Parse("yaml/example.yaml")
	if err != nil {
		cly.Error.Println(err)
	}
	err = parse(config, os.Args)
	if err != nil {
		cly.Error.Println(err)
	}
}
