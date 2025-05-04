# ⛩️ cly

[![Go Report Card](https://goreportcard.com/badge/github.com/TimoKats/cly)](https://goreportcard.com/report/github.com/TimoKats/cly)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![GitHub tag](https://img.shields.io/github/tag/TimoKats/cly?include_prereleases=&sort=semver&color=blue)](https://github.com/TimoKats/cly/releases/)

Cly allows you to create feature rich aliases in YAML. The example below is an example YAML with 3 aliases (update, dashboard, test) with different configuration options. You can add this file in `~/.cly.yaml` or somewhere custom by setting the env `CLYPATH`.  

Cly has two commands: `cly run <<command>>` and `cly ls`.  You can install cly using: `go install github.com/TimoKats/cly@latest`.

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
  
## Docs
This table shows an overview of the fields that can be supplied in your YAML alias objects.

<table>
  <thead>
    <tr>
      <th width="500px">Field</th>
      <th width="500px">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr width="600px">
      <td>command/commands</td>
      <td>The alias command. Can be a list of commands or one command.</td>
    </tr>
    <tr width="600px">
      <td>name</td>
      <td>Name of the alias. Mandatory for subcommands. Root commands derive the name from the YAML name (see above).</td>
    </tr>
    <tr width="600px">
      <td>dir</td>
      <td>Directory to run the alias in. If empty, current working directory.</td>
    </tr>
    <tr width="600px">
      <td>envs</td>
      <td>Add additional env variables for the alias. List of name/value pairs.</td>
    </tr>
    <tr width="600px">
      <td>concurrent</td>
      <td>Boolean. If true (and multiple commands are supplied), then the commands are executed in concurrent threads.</td>
    </tr>
    <tr width="600px">
      <td>subcommands</td>
      <td>List of command objects (i.e. all other fields apply) that become subcommands. E.g. cly run *command* *subcommand*</td>
    </tr>
  </tbody>
</table>

&nbsp;

You can pass parameters to your aliases when invoking them. For this, we use bash syntax. Adding `$@` adds all parameters to an alias. `$0...n` inserts an alias based on the index. The example above has some examples for this.


