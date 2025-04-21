package main

import (
	"errors"

	cly "github.com/TimoKats/cly/internal"

	"os"
)

func run(config cly.Config, args []string) error {
	config.AddArgs(args)
	if len(args) < 3 {
		return errors.New("no command to run provided")
	}
	return config.RunAlias(args[2])
}

func parse(config cly.Config, args []string) error {
	switch {
	case len(args) == 1:
		return errors.New("no command found")
	case args[1] == "run":
		return run(config, args)
	case args[1] == "ls":
		return config.List()
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
