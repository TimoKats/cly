// config (i.e. the yaml with your aliases) related functions. Often calls
// modules from alias and is called by the main control flow.

package internal

import (
	"os"

	"gopkg.in/yaml.v2"
)

func (config *Config) GetAlias(args []string, aliasIndex int) (*Alias, bool) {
	var aliasName string = args[aliasIndex]
	alias, ok := config.aliases[aliasName]
	for aliasIndex < len(args)-1 && ok {
		aliasIndex += 1
		aliasName = args[aliasIndex]
		if match := find(alias.subAliases(), aliasName); match != -1 {
			alias = alias.Subs[match]
		} else {
			aliasIndex -= 1
			break
		}
	}
	return alias, alias != nil
}

func (config *Config) List(tree bool) error {
	for name, alias := range config.aliases {
		if tree {
			for index := range alias.Subs {
				alias.Subs[index].print("-") // recursive
			}
		} else {
			printTable([]string{name, alias.Command})
		}
	}
	return nil
}

func (config *Config) AddArgs(args []string) {
	if len(args) <= 3 { // no args provided
		return
	}
	for _, alias := range config.aliases {
		popIndex := 3
		alias.Args = args[popIndex:]
		for _, alias := range alias.Subs {
			alias.addArgs(args, popIndex+1) // recursive
		}
	}
}

func Parse(filename string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	// unmarshal the YAML data into the map
	err = yaml.Unmarshal(data, &config.aliases)
	return config, err
}
