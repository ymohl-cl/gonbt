# GONBT

Go because this project is written in Golang and NBT because it is a nbt parser &nbsp;:v:&nbsp;

![Build status](https://github.com/ymohl-cl/gonbt/actions/workflows/go.yml/badge.svg)
[![Codecov](https://codecov.io/gh/ymohl-cl/gonbt/branch/main/graph/badge.svg?token=BD09YAUX00&)](https://codecov.io/gh/ymohl-cl/gonbt)

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/ymohl-cl/gonbt)
[![Sourcegraph](https://sourcegraph.com/github.com/ymohl-cl/gonbt/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/ymohl-cl/gonbt?badge)
[![GoReportCard](https://goreportcard.com/badge/github.com/ymohl-cl/gonbt?style=flat-square)](https://goreportcard.com/report/github.com/ymohl-cl/gonbt)
[![Discord](https://img.shields.io/badge/Discord-%40gonbt-informational?style=flat-square)](https://discord.gg/NQEenAqBEe)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gonbt/main/LICENSE)

- [GONBT](#gonbt)
  - [Description](#description)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Roadmap](#roadmap)
  - [Contributing](#contributing)
  - [Licence](#licence)

## Description

The __Named Binary Tag__ (NBT) format is used by Minecraft for the various files in which it saves data... the next on [__wiki__](https://minecraft.fandom.com/wiki/NBT_format) &nbsp;:sunglasses:&nbsp;

This gonbt package marshal and unmarshal nbt data in a readable format [see tag.go](https://github.com/ymohl-cl/gonbt/blob/main/tag.go)

## Installation

Easy to install like all go packages &nbsp;:grin:&nbsp;

``` bash
<$ go get -u github.com/ymohl-cl/gonbt
```

## Usage

``` Golang
// To read
func main() {
    var dataIn []byte // your nbt data
    var err error

    if tag, err = gonbt.Unmarshal(dataIn); err != nil {
      panic(err)
    }
}
```

``` Golang
// To write
func main() {
    var tag gonbt.Tag // your data to convert in nbt format
    var err error

    if tag, err = gonbt.Marshal(tag, gonbt.CompressNone); err != nil {
      panic(err)
    }
}
```

## Roadmap

The next should be convert Tag to golang struct with specific field's tag like json or other parser stuf

## Contributing

&nbsp;:grey_exclamation:&nbsp; Use issues for everything

- For a small change, just send a PR.
- For bigger changes open an issue for discussion before sending a PR.
- PR should have:
  - Test case
  - Documentation
  - Example (If it makes sense)
- You can also contribute by:
  - Reporting issues
  - Suggesting new features or enhancements
  - Improve/fix documentation

Thank you &nbsp;:pray:&nbsp;&nbsp;:+1:&nbsp;

## Licence

[MIT](https://github.com/ymohl-cl/gonbt/blob/main/LICENSE)
