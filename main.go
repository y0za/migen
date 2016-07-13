package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Print(`usage: migen <command> [option]

Commands:
	new [options] : Generate a migration file
`)
		os.Exit(0)
	}
}

func newCommand(args []string) {
	var (
		format string
		name   string
	)
	fs := flag.NewFlagSet("new", flag.ExitOnError)
	fs.StringVar(&format, "format", "date", "specify the format of prefix")
	fs.StringVar(&name, "name", "", "specify the file name")

	fs.Parse(args)

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

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}
	switch args[0] {
	case "new":
		newCommand(args[1:])
	default:
		flag.Usage()
	}
}
