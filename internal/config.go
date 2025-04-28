package internal

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// Returns the alias based on the provided args (bool returned for success)
func (config *Config) GetAlias(args []string, aliasIndex int) (*Alias, bool) {
	var aliasName string = args[aliasIndex]
	alias, ok := config.aliases[aliasName]
	for aliasIndex < len(args)-1 && ok {
		aliasIndex += 1
		aliasName = args[aliasIndex]
		if match := find(alias.subAliases(), aliasName); match != -1 {
			alias = alias.Subs[match]
		} else {
			aliasIndex -= 1 //nolint
			break
		}
	}
	return alias, alias != nil
}

// Lists aliases, args can contain alias to list subcommands
func (config *Config) List(args []string) error {
	if len(args) > 2 { // NOTE: keep this?
		if alias, ok := config.aliases[args[2]]; ok {
			for _, subAlias := range alias.Subs {
				printTable([]string{subAlias.Name, subAlias.Command})
			}
			return nil
		}
		return errors.New("alias '" + args[2] + "' not found")
	}
	for name, alias := range config.aliases {
		printTable([]string{name, alias.Command})
	}
	return nil
}

// Fills in the provided arguments for variables ($@, $0)
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

// Reads the cly yaml file and returns a config object
func Parse() (Config, error) {
	var path string = configPath()
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	// unmarshal the YAML data into the map
	err = yaml.Unmarshal(data, &config.aliases)
	return config, err
}
