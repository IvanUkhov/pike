package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	ivy = flag.String("ivy", "ivy", "the path to ivy")
)

func main() {
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: pike [flags]\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(2)
}
