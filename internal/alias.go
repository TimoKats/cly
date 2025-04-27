package internal

import (
	"os/exec"
	"strings"
)

// Puts the provided args in placeholders ($@, $0) of the alias.
func (alias *Alias) formatCommand() (string, []string) {
	var cmd []string = strings.Fields(alias.Command)
	for index, value := range cmd {
		if value == "$@" {
			cmd = popIndex(cmd, index)
			cmd = append(cmd, alias.Args...)
		}
		if argIndex, ok := getParam(value); ok {
			if len(alias.Args) <= argIndex {
				cmd = popIndex(cmd, index)
			} else {
				cmd[index] = alias.Args[argIndex]
			}
		}
	}
	return cmd[0], cmd[1:]
}

// Adds the command args to the alias object.
func (alias *Alias) addArgs(args []string, popIndex int) {
	if len(args) <= popIndex {
		return
	}
	alias.Args = args[popIndex:]
	for index := range alias.Subs {
		alias.Subs[index].addArgs(args, popIndex+1)
	}
}

// Returns the names for sub aliases in a string list.
func (alias *Alias) subAliases() []string {
	var subAliases []string
	for _, alias := range alias.Subs {
		subAliases = append(subAliases, alias.Name)
	}
	return subAliases
}

// Runs the alias and returns command error.
func (alias *Alias) Run() error {
	app, args := alias.formatCommand()
	cmd := exec.Command(app, args...)
	cmd.Dir = alias.Dir
	output, err := cmd.CombinedOutput()
	Info.Println(string(output))
	return err
}
