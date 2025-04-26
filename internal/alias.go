// alias related functions. Typically called by the config module.

package internal

import (
	"os/exec"
	"strings"
)

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

func (alias *Alias) addArgs(args []string, popIndex int) {
	if len(args) <= popIndex {
		return
	}
	alias.Args = args[popIndex:]
	for index := range alias.Subs {
		alias.Subs[index].addArgs(args, popIndex+1)
	}
}

func (alias *Alias) print(indent string) {
	Info.Printf("%s %s: %s", indent, alias.Name, alias.Command)
	for index := range alias.Subs {
		alias.Subs[index].print("  " + indent)
	}
}

func (alias *Alias) subAliases() []string {
	var subAliases []string
	for _, alias := range alias.Subs {
		subAliases = append(subAliases, alias.Name)
	}
	return subAliases
}

func (alias *Alias) Run() error {
	app, args := alias.formatCommand()
	cmd := exec.Command(app, args...)
	cmd.Dir = alias.Dir
	output, err := cmd.CombinedOutput()
	Info.Println(string(output))
	return err
}
