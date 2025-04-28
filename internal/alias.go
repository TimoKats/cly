package internal

import (
	"os/exec"
	"strings"
)

// Puts the provided args in placeholders ($@, $0) of the alias.
// Update: command is provided as arg, because I want to reuse it for lists.
func (alias *Alias) formatCommand(command string) (string, []string) {
	var cmd []string = strings.Fields(command)
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
	var output []byte
	var err error
	for _, command := range append(alias.Commands, alias.Command) {
		// skip empty commands (breaks formatting)
		if len(command) == 0 {
			continue
		}
		// setup
		app, args := alias.formatCommand(command)
		cmd := exec.Command(app, args...)
		cmd.Dir = alias.Dir
		// exec
		output, err = cmd.CombinedOutput()
		Info.Println(string(output))
	}
	return err
}
