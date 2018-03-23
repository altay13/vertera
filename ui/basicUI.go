package ui

import (
	"fmt"
	"io"
	"time"
)

type BasicUI struct {
	Reader io.Reader
	Writer io.Writer
}

// TODO: implement later
func (ui *BasicUI) Ask(msg string) (string, error) {
	fmt.Fprintf(ui.Writer, "%s", msg)
	return "", nil
}

func (ui *BasicUI) Output(msg string) {
	fmt.Fprintf(ui.Writer, "%s >>> %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (ui *BasicUI) Info(msg string) {
	fmt.Fprintf(ui.Writer, "%s >>> %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (ui *BasicUI) Error(msg string) {
	fmt.Fprintf(ui.Writer, "%s >>> %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}
