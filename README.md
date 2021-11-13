<h1 align="center">
  <br>xlsx2mysql<br>
</h1>

<h4 align="center">An tool of helping your fastly generate SQL from Excel</h4>

#### [中文文档](https://github.com/ToWeLong/xlsx2mysql/blob/master/README_zh-CN.md)

## Origin

In order to convert Excel to MySQL and I made a tool to implement.But Why do I write it? because I find a round tools that it still not satisfy my needs.Then it was born.

## Install

### Install from source

xlsx2mysql Requires Go >= 1.16. You can build it from source:

```sh
$ go get -u -v github.com/ToWeLong/xlsx2mysql
```

### Install pre-build binariy

Pre-built binaries are available here: [release](https://github.com/ToWeLong/xlsx2mysql/releases)

## Usage

### Generate to sql

```
$ xlsx2mysql generate example.xlsx
```


## Interface

### Help

```
An tool of helping your fastly generate SQL from Excel.

Usage:
  xlsx2mysql [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  generate    generate SQL from Excel
  help        Help about any command

Flags:
  -h, --help   help for xlsx2mysql

Use "xlsx2mysql [command] --help" for more information about a command.
```

## License

MIT