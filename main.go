package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/peterh/liner"
)

func main() {
	reader, writer := io.Pipe()

	liner := liner.NewLiner()
	liner.SetCtrlCAborts(true)
	defer liner.Close()

	cmd := exec.Command("ivy", os.Args[1:]...)
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
