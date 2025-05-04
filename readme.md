# ⛩️ cly

[![Go Report Card](https://goreportcard.com/badge/github.com/TimoKats/cly)](https://goreportcard.com/report/github.com/TimoKats/cly)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![GitHub tag](https://img.shields.io/github/tag/TimoKats/cly?include_prereleases=&sort=semver&color=blue)](https://github.com/TimoKats/cly/releases/)

Use cly to define feature-rich aliases in YAML. Cly has two commands: `cly run <<command>>` runs an alias and `cly ls/tree` lists the current aliases in your yaml.  

In your YAML, each object is an alias with potential subcommands, directories, and parameters. The yaml below shows the configuration options in practice.

```yaml

update:
  command: /some/path/script.sh $@  # adds args to alias. E.g.: cly run update <<x,y,z>>
  subcommands:
  - name: ping  # subcommand for alias, called with: cly run update <<ping>>
    command: /some/other/path/script.sh

dashboard:
  command: streamlit run main.py
  dir: /path/to/python/  # sets a directory to run an alias in

test:
  command: $0 test.py  # Insert args based on index. E.g.: cly run test <<python3.12>>
                       # Runs <<python3.12>> main.py

```
  
&nbsp;

You can install cly using: `go install github.com/TimoKats/cly@latest`. You can add your aliases in your home directory `~/.cly.yaml`. Should work on all operating systems.
