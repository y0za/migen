# migen

[![Build Status](https://travis-ci.org/y0za/migen.svg?branch=master)](https://travis-ci.org/y0za/migen) [![Go Report Card](https://goreportcard.com/badge/github.com/y0za/migen)](https://goreportcard.com/report/github.com/y0za/migen)

template generator tool for [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)

## Installation
```
$ go get -u github.com/y0za/migen
```

## Usage
```
usage: migen new [option]

Options:
	-f, --format <format> specify the format of file prefix
	                      formats: date, counter, unix, none (Default: date)
	-n, --name <name>     specify the file name
	-h, --help            show this help
```
Each format example (when name is `"foo"`)
- date: `20160714234637foo.sql`
- counter: `1foo.sql` (count up from existing files)
- unix: `1468507767foo.sql` (unix timestamp)
- none: `foo.sql`

## License
MIT License

## Author
yoza (y0za)
