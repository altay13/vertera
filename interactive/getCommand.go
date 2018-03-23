package interactive

import (
	"fmt"

	"github.com/altay13/vertera/eventHandler"
	"github.com/altay13/vertera/ui"
)

type GET struct {
	args []string

	handler *eventHandler.EventHandler
	ui      *ui.InteractiveUI
}

func GetCommand(args []string, handler *eventHandler.EventHandler, ui *ui.InteractiveUI) InterCMD {
	return &GET{
		args:    args,
		handler: handler,
		ui:      ui,
	}
}

func (cmd *GET) Validate() error {
	if len(cmd.handler.GetDBName()) == 0 {
		return fmt.Errorf("Please connect to database. Type `help` if you don't know what to do.")
	}
	return nil
}

func (cmd *GET) Run() {

	if err := cmd.Validate(); err != nil {
		cmd.ui.Error(err.Error())
		return
	}

	for _, v := range cmd.args {
		respCh := make(chan *eventHandler.Event)

		request := eventHandler.Request{
			Cmd:      eventHandler.GET,
			Response: respCh,
			Event: &eventHandler.Event{
				Key: v,
			},
		}

		cmd.handler.RequestChan <- request

		select {
		case resp := <-respCh:
			if resp.Err != nil {
				cmd.ui.Error(resp.Err.Error())
			} else {
				cmd.ui.Output(fmt.Sprintf("Key = %s", resp.Key))
				cmd.ui.Output(fmt.Sprintf("Value = %s", resp.Value))
			}
		}
		// TODO: add timeout exception
	}
}
