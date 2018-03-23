package ui

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type UIColor struct {
	Code int
}

var (
	UIColorNone      UIColor = UIColor{-1}
	UIColorRed               = UIColor{31}
	UIColorGreen             = UIColor{32}
	UIColorYellow            = UIColor{33}
	UIColorBlue              = UIColor{34}
	UIColorLightCyan         = UIColor{96}
)

type InteractiveUI struct {
	Ui *BasicUI

	OutputColor UIColor
	InfoColor   UIColor
	ErrorColor  UIColor
	AskColor    UIColor

	respCh chan string
	asked  bool
}

func DefaultInteractiveUI() *InteractiveUI {
	ui := &InteractiveUI{
		Ui: &BasicUI{
			Reader: bufio.NewReader(os.Stdin),
			Writer: os.Stdout,
		},
		OutputColor: UIColorGreen,
		InfoColor:   UIColorYellow,
		ErrorColor:  UIColorRed,
		AskColor:    UIColorLightCyan,
		asked:       false,
		respCh:      make(chan string, 1),
	}

	return ui
}

func (ui *InteractiveUI) StartInteractive(cmdCh chan<- string) {
	scanner := bufio.NewScanner(ui.Ui.Reader)

	for scanner.Scan() {
		cmd := scanner.Text()
		if ui.asked {
			ui.respCh <- cmd
			continue
		}
		cmdCh <- cmd
	}
}

func (ui *InteractiveUI) Ask(msg string) (string, error) {
	ui.asked = true

	ui.Ui.Output(ui.fill(msg, ui.AskColor))

	timer := time.NewTimer(60 * time.Second)

	select {
	case cmd := <-ui.respCh:
		ui.asked = false
		return cmd, nil
	case <-timer.C:
		ui.asked = false
		return "", fmt.Errorf("Interrupted. Timeout for request")
	}

	return "", nil
}

func (ui *InteractiveUI) Output(msg string) {
	ui.Ui.Output(ui.fill(msg, ui.OutputColor))
}

func (ui *InteractiveUI) Info(msg string) {
	ui.Ui.Output(ui.fill(msg, ui.InfoColor))
}

func (ui *InteractiveUI) Error(msg string) {
	ui.Ui.Output(ui.fill(msg, ui.ErrorColor))
}

func (ui *InteractiveUI) fill(msg string, color UIColor) string {
	return fmt.Sprintf("\033[0;%dm%s\033[0m", color.Code, msg)
}
