package internal

type Alias struct {
	Command string `yaml:"command"`

	// optional
	Subs []*Alias `yaml:"subcommands"`
	Name string   `yaml:"name"`
	Dir  string   `yaml:"dir"`
	Args []string
}

type Config struct {
	aliases map[string]*Alias
}
