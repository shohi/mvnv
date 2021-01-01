# mvnv
maven version manager written in go, inspired by helmevn version manager - [helmenv](https://github.com/little-angry-clouds/kubernetes-binaries-managers/tree/master/cmd/helmenv) and golang version manager - [g](https://github.com/voidint/g).

## Install

```bash
make install
```

Note: make sure `$GOBIN` is in `$PATH`

## Usage

```bash
$> mvnv --help

maven version manager

Usage:
  mvnv [command]

Available Commands:
  clean       Remove downloaded source file
  help        Help about any command
  install     Install a specific version
  list        List all installed versions
  list-remote List all installable versions
  uninstall   Uninstall a specific version
  use         Switch to specific version
  version     Print version

Flags:
  -h, --help              help for mvnv
      --loglevel string   log level (default "INFO")

Use "mvnv [command] --help" for more information about a command.

```

## Links
1. voidint/g, https://github.com/voidint/g
1. little-angry-clouds/kubernetes-binaries-managers, <https://github.com/little-angry-clouds/kubernetes-binaries-managers>
