# ⛩️ cly

Use cly to define feature-rich aliases in YAML. For example, you can pass parameters, directories, or define sub commands.  

Cly has two commands: `cly run <<command>>` runs an alias and `cly ls/tree` lists the current aliases in your yaml. For configuration, the YAML below shows most available functionalities.

```yaml

update:
  command: /some/path/script.sh $@  # adds args to your alias. E.g.: cly run update <<something>>
  subcommands:
  - name: ping  # subcommand, called with: cly run update ping
    command: /some/other/path/script.sh

dashboard:
  command: streamlit run main.py
  dir: /path/to/python/  # set a directory to run an alias in

```

You can install cly using: `go install github.com/TimoKats/cly@latest`. You can add your aliases in your home directory `~/.cly.yaml`. Should work on all operating systems.