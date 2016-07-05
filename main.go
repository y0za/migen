package main

import (
	"flag"
)

var (
	showHelp      bool
	prefixFormat  string
	migrationName string
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "shorthand of help")
	flag.BoolVar(&showHelp, "help", false, "show this help")
	flag.StringVar(&prefixFormat, "f", "date", "shorthand of format")
	flag.StringVar(&prefixFormat, "format", "date", "specify the file prefix format")
	flag.StringVar(&migrationName, "n", "", "shorthand of name")
	flag.StringVar(&migrationName, "name", "", "specify the file name")
}

func main() {
	flag.Parse()

	if showHelp {
		flag.Usage()
	}
}
