# Runic CLI

A basic CLI wrapper around [the Runic parser library](https://github.com/vqvw/runic).

## Installation

```sh
go get github.com/vqvw/runic_cli
cd $GOPATH/src/github.com/vqvw/runic_cli
go install
```

## Usage

- Outputs parsed tree as JSON to stdout.

```sh
runic_cli -in my_document.runic
```

- Outputs parsed tree as JSON to file.

```sh
runic_cli -in my_document.runic -out my_document.json
```

- Outputs parsed tree as HTML to file.

```sh
runic_cli -in my_document.runic -out my_document.html
```
