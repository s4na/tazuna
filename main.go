package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed VERSION
var versionString string

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version (long form)")
	flag.Parse()

	if showVersion {
		fmt.Println(strings.TrimSpace(versionString))
		return
	}
	fmt.Println("hello world")
}
