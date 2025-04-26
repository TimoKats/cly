// Types of cly. A config refers to the yaml that's loaded (and consists of
// aliases). The alias attribute Args is not derived from YAML but instead
// filled in later.

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
