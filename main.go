package main

import (
	"flag"
	"fmt"
	"os"
)

func newCommand(args []string) {
	var (
		format string
		name   string
	)
	fs := flag.NewFlagSet("new", flag.ExitOnError)
	fs.StringVar(&format, "format", "date", "shorthand of format")
	fs.StringVar(&name, "name", "", "specify the file name")

	fs.Parse(args)
	fmt.Println("format:", format)
	fmt.Println("name:", name)
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		os.Exit(0)
	}
	switch args[0] {
	case "new":
		newCommand(args[1:])
	default:
		os.Exit(0)
	}
}
