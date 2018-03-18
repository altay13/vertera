package interactive

import (
	"fmt"

	"github.com/altay13/vertera/eventHandler"
)

type GET struct {
	args []string

	handler *eventHandler.EventHandler
}

func GetCommand(args []string, handler *eventHandler.EventHandler) InterCMD {
	return &GET{
		args:    args,
		handler: handler,
	}
}

func (cmd *GET) Validate() error {
	return nil
}

func (cmd *GET) Run() string {

	if err := cmd.Validate(); err != nil {
		return err.Error()
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
				return resp.Err.Error()
			}
			return fmt.Sprintf("\n%s\n", string(resp.Value))
		}

		// TODO: create an object and send it to SET routine for saving into DB!
	}

	return "get command is performed"
}
