# g2lint

Parse and validate server and evasion profiles for the Gambit-Framework. It is recommended to validate any server or evasion profiles before loading them.

## Quick Start

### Build From Source

```shell
git clone git@github.com:Gambit-Framework/g2lint.git
cd g2lint
make
```

### Running

```shell
build/g2lint --help
usage: g2lint [-h|--help] [-s|--server-profile "<value>"] [-e|--evasion-profile
              "<value>"] [-v|--verbose]

              parse and validate server and evasion profile

Arguments:

  -h  --help             Print help information
  -s  --server-profile   path to server profile
  -e  --evasion-profile  path to evasion profile
  -v  --verbose          show parsed data. Default: false
```

> `g2lint` needs to run on the same host the teamserver will run on

#### Validate Server Profile

```shell
build/g2lint --server-profile path/to/profile
```

#### Validate Evasion Profile

```shell
build/g2lint --evasion-profile path/to/profile
```

## Known Errors

- Due to the way the kdl-go library parses kdl, if two listeners have the same name, only the second listener will be parsed
