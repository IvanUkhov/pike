package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/peterh/liner"
)

var (
	path = flag.String("ivy", "ivy", "the path to ivy")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	reader, writer := io.Pipe()

	liner := liner.NewLiner()
	liner.SetCtrlCAborts(true)
	defer liner.Close()

	cmd := exec.Command(*path)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = reader, os.Stdout, os.Stderr
	cmd.Start()

	for {
		if line, err := liner.Prompt(""); err == nil {
			liner.AppendHistory(line)
			fmt.Fprintln(writer, line)
		} else {
			break
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: pike [flags]\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(2)
}
