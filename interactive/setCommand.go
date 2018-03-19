package interactive

import (
	"fmt"
	"strings"

	"github.com/altay13/vertera/eventHandler"
	"github.com/altay13/vertera/ui"
)

type SET struct {
	args []string

	handler *eventHandler.EventHandler
	ui      *ui.InteractiveUI
}

func SetCommand(args []string, handler *eventHandler.EventHandler, ui *ui.InteractiveUI) InterCMD {
	return &SET{
		args:    args,
		handler: handler,
		ui:      ui,
	}
}

func (cmd *SET) Validate() error {
	if !strings.ContainsAny(strings.Join(cmd.args, " "), "=") {
		return fmt.Errorf("Syntax error in command. Please enter help for manual.")
	}
	return nil
}

func (cmd *SET) Run() {
	if err := cmd.Validate(); err != nil {
		cmd.ui.Error(err.Error())
		return
	}

	for _, v := range cmd.args {
		dt := strings.Split(v, "=")

		respCh := make(chan *eventHandler.Event)
		request := eventHandler.Request{
			Cmd:      eventHandler.SET,
			Response: respCh,
			Event: &eventHandler.Event{
				Key:   dt[0],
				Value: []byte(dt[1]),
			},
		}
		cmd.handler.RequestChan <- request
		select {
		case resp := <-respCh:
			if resp.Err != nil {
				cmd.ui.Error(fmt.Sprintf("Failed to set the Key = %s.", resp.Key))
				cmd.ui.Error(resp.Err.Error())
				continue
			} else {
				cmd.ui.Output(fmt.Sprintf("The Key = %s is set.", resp.Key))
				cmd.ui.Output(fmt.Sprintf("Value = %s", resp.Value))
			}
		}

		// TODO: create an object and send it to SET routine for saving into DB!
	}
}
