package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed VERSION
var version string

func main() {
	if len(os.Args) >= 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Println(strings.TrimSpace(version))
		return
	}
	fmt.Println("hello world")
}
