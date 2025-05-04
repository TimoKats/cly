package internal

// Read from YAML. Args is added based on command.
type Alias struct {
	// at least one needs to be set
	Command  string   `yaml:"command"`
	Commands []string `yaml:"commands"`

	Subs       []*Alias `yaml:"subcommands"`
	Name       string   `yaml:"name"`
	Dir        string   `yaml:"dir"`
	Concurrent bool     `yaml:"concurrent"`
	Args       []string
}

// YAML file is read into this struct. List of alias structs.
type Config struct {
	aliases map[string]*Alias
}
