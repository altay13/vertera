package ui

import (
	"bufio"
	"fmt"
	"os"
)

type UIColor struct {
	Code int
}

var (
	UIColorNone   UIColor = UIColor{-1}
	UIColorRed            = UIColor{31}
	UIColorGreen          = UIColor{32}
	UIColorYellow         = UIColor{33}
	UIColorBlue           = UIColor{34}
)

type InteractiveUI struct {
	Ui *BasicUI

	OutputColor UIColor
	InfoColor   UIColor
	ErrorColor  UIColor
}

func DefaultInteractiveUI() *InteractiveUI {
	ini := &InteractiveUI{
		Ui: &BasicUI{
			Reader: bufio.NewReader(os.Stdin),
			Writer: os.Stdout,
		},
		OutputColor: UIColorGreen,
		InfoColor:   UIColorYellow,
		ErrorColor:  UIColorRed,
	}

	return ini
}

func (ini *InteractiveUI) Output(msg string) {
	ini.Ui.Output(ini.fill(msg, ini.OutputColor))
}

func (ini *InteractiveUI) Info(msg string) {
	ini.Ui.Output(ini.fill(msg, ini.InfoColor))
}

func (ini *InteractiveUI) Error(msg string) {
	ini.Ui.Output(ini.fill(msg, ini.ErrorColor))
}

func (ini *InteractiveUI) fill(msg string, color UIColor) string {
	return fmt.Sprintf("\033[0;%dm%s\033[0m", color.Code, msg)
}
