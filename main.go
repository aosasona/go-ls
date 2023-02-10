package main

import (
	"flag"
	"fmt"
	"go-ls/cmd/gols"
	"os"
)

func main() {
	var path string
	showHidden := flag.Bool("a", false, "Show hidden files")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		path = string(args[0])
	} else {
		path = "."
	}

	if err := gols.Run(gols.Config{
		ShowHidden: *showHidden,
		Path:       path,
	}); err != nil {
		fmt.Printf("error: %s", err.Error())
		os.Exit(1)
	}
}
