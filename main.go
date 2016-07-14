package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	mainUsage = `usage: migen <command> [option]

Commands:
	new [options] : Generate a migration file`

	newUsage = `usage: migen new [option]

Options:
	-f, --format <format> specify the format of file prefix
	                      formats: date, counter, unix, none (Default: date)
	-n, --name <name>     specify the file name
	-h, --help            show this help`
)

func init() {
	flag.Usage = func() {
		fmt.Println(mainUsage)
	}
}

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "new":
		newCommand(flag.Args()[1:])
	default:
		flag.Usage()
	}
}

func newCommand(args []string) {
	var (
		format string
		name   string
		help   bool
	)
	fs := flag.NewFlagSet("new", flag.ExitOnError)
	fs.StringVar(&format, "format", "date", "specify the format of prefix")
	fs.StringVar(&format, "f", "date", "shorthand of format")
	fs.StringVar(&name, "name", "", "specify the file name")
	fs.StringVar(&name, "n", "", "shorthand of name")
	fs.BoolVar(&help, "help", false, "show help")
	fs.BoolVar(&help, "h", false, "shorthand of help")

	fs.Parse(args)

	if help {
		fmt.Println(newUsage)
		return
	}

	fn, err := fileName(format, name)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = createFile(fn, ".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fn, "is generated")
}
