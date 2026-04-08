package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

//go:embed VERSION
var versionString string

func run(args []string, out io.Writer, errOut io.Writer) int {
	fs := flag.NewFlagSet("tazuna", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var showVersion bool
	fs.BoolVar(&showVersion, "v", false, "show version")
	fs.BoolVar(&showVersion, "version", false, "show version (long form)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if showVersion {
		fmt.Fprintln(out, strings.TrimSpace(versionString))
		return 0
	}
	fmt.Fprintln(out, "hello world")
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
