package internal

import (
	"os/exec"
	"strconv"
	"strings"
)

// this makes me dislike golang...
func popIndex(values []string, delIndex int) []string {
	var new []string
	for index, value := range values {
		if index != delIndex {
			new = append(new, value)
		}
	}
	return new
}

// return true if it's a param, and the index of the param
func getParam(value string) (int, bool) {
	var isParam bool
	var index int64
	if len(value) > 1 && value[0] == '$' {
		isParam = true
		index, _ = strconv.ParseInt(value[1:], 10, 0)
	}
	return int(index), isParam
}

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

func (alias *Alias) run() error {
	app, args := alias.formatCommand()
	cmd := exec.Command(app, args...)
	cmd.Dir = alias.Dir
	output, err := cmd.CombinedOutput()
	Info.Println(string(output))
	return err
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
