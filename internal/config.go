package internal

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

func (config *Config) RunAlias(aliasName string) error {
	if alias, ok := config.aliases[aliasName]; ok {
		return alias.run()
	}
	return errors.New("alias '" + aliasName + "' not found in config")
}

func (config *Config) List() error {
	for name, alias := range config.aliases {
		Info.Printf("%s: %s", name, alias.Command)
		for index := range alias.Subs {
			alias.Subs[index].print("  >") // recursive
		}
	}
	return nil
}

func (config *Config) AddArgs(args []string) {
	if len(args) <= 3 {
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
	// Unmarshal the YAML data into the map
	err = yaml.Unmarshal(data, &config.aliases)
	return config, err
}
