package internal

import (
	"os"
	"os/exec"
	"strings"
	"sync"
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

// Takes envs from yaml and creates a .env file (kinda) for exec library.
func (alias *Alias) formatEnv() []string {
	var envs []string
	for _, env := range alias.Envs {
		envs = append(envs, env.Name+"="+env.Value)
	}
	return envs
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

// executes a command, logs the result, and returns error.
func (alias *Alias) exec(command string, wg *sync.WaitGroup) error {
	// skip empty commands (breaks formatting)
	if len(command) == 0 {
		return nil
	}
	// only for concurrent runs
	if wg != nil {
		defer wg.Done()
	}
	// setup
	app, args := alias.formatCommand(command)
	cmd := exec.Command(app, args...)
	cmd.Env = append(os.Environ(), alias.formatEnv()...)
	cmd.Dir = alias.Dir
	// exec
	output, err := cmd.CombinedOutput()
	Info.Println(string(output))
	return err
}

// Runs the alias in sequential order.
func (alias *Alias) SequentialRun() error {
	var err error
	for _, command := range append(alias.Commands, alias.Command) {
		err = alias.exec(command, nil)
	}
	return err
}

// Runs the alias commands concurrently. TODO: use errors?
func (alias *Alias) ConcurrentRun() error {
	var wg sync.WaitGroup
	var commands []string = append(alias.Commands, alias.Command)
	var workers = len(commands) - 1
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go alias.exec(commands[i], &wg) //nolint
	}
	wg.Wait()
	Info.Printf("Finished %d runs concurrently.", workers)
	return nil
}
