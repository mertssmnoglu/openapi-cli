# OpenAPI CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/mertssmnoglu/openapi-cli)](https://goreportcard.com/report/github.com/mertssmnoglu/openapi-cli)
[![GoDoc](https://godoc.org/github.com/mertssmnoglu/openapi-cli?status.svg)](https://godoc.org/github.com/mertssmnoglu/openapi-cli)
![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mertssmnoglu/openapi-cli)

A Command Line Interface (CLI) tool for serving, validating, and linting OpenAPI specifications.

## Table of Contents

- [Installation](#installation)
    - [GitHub Releases](#github-releases)
    - [Command Line](#command-line)
- [Usage](#usage)
- [License](#license)

## Installation

### GitHub Releases

Please visit the [GitHub Releases](https://github.com/mertssmnoglu/openapi-cli/releases) page to download the latest version of the OpenAPI CLI.

### Command Line

To install the latest version of the OpenAPI CLI via the command line, you can use the following command:

```shell
go install github.com/mertssmnoglu/openapi-cli@latest
```

It will automatically download the latest version of the OpenAPI CLI and install it in your `$GOPATH/bin` directory. Make sure that your `$GOPATH/bin` directory is in your `$PATH`.

## Usage

```shell
>>> openapi-cli help

A Command Line Interface (CLI) tool for serving, validating, and linting OpenAPI specifications.

Usage:
  openapi-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  serve       Serve the OpenAPI file

Flags:
  -h, --help   help for openapi-cli

Use "openapi-cli [command] --help" for more information about a command.
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
