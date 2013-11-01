//Package gnuplot provides convenience wrappers for piping to gnuplot.
package gnuplot

import (
	"fmt"
	"io"
	"os/exec"
)

type Plot chan string

func handleGnuplot(p Plot, pipe io.WriteCloser) {
	var command string
	ok := true

	for ok {
		command, ok = <-p
		if !ok {
			pipe.Close()
			continue
		}
		fmt.Fprintln(pipe, command)
	}
}

func New(persist bool) (Plot, error) {
	var cmd *exec.Cmd
	if persist {
		cmd = exec.Command("gnuplot", "-persist")
	} else {
		cmd = exec.Command("gnuplot")
	}
	pipe, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	p := make(chan string)
	go handleGnuplot(p, pipe)
	return p, nil
}
