package internal

// Read from YAML. Args is added based on command.
type Alias struct {
	Command string `yaml:"command"`

	// optional
	Subs []*Alias `yaml:"subcommands"`
	Name string   `yaml:"name"`
	Dir  string   `yaml:"dir"`
	Args []string
}

// YAML file is read into this struct. List of aliases.
type Config struct {
	aliases map[string]*Alias
}
