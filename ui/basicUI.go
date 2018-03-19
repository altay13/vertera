package ui

import (
	"fmt"
	"io"
)

type BasicUI struct {
	Reader io.Reader
	Writer io.Writer
}

func (ui *BasicUI) Output(msg string) {
	fmt.Fprintf(ui.Writer, ">>> %s\n", msg)
}

func (ui *BasicUI) Info(msg string) {
	fmt.Fprintf(ui.Writer, ">>> %s\n", msg)
}

func (ui *BasicUI) Error(msg string) {
	fmt.Fprintf(ui.Writer, ">>> %s\n", msg)
}
