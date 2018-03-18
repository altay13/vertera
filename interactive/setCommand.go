package interactive

import (
	"fmt"
	"strings"

	"github.com/altay13/vertera/eventHandler"
)

type SET struct {
	args []string

	handler *eventHandler.EventHandler
}

func SetCommand(args []string, handler *eventHandler.EventHandler) InterCMD {
	return &SET{
		args:    args,
		handler: handler,
	}
}

func (cmd *SET) Validate() error {
	if !strings.ContainsAny(strings.Join(cmd.args, " "), "=") {
		return fmt.Errorf("Syntax error in command. Please enter help for manual.")
	}
	return nil
}

func (cmd *SET) Run() string {
	respStr := ""

	if err := cmd.Validate(); err != nil {
		return err.Error()
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
			respStr = fmt.Sprintf("%s\n%s", respStr, resp.Err.Error())
		}

		// TODO: create an object and send it to SET routine for saving into DB!
	}

	return fmt.Sprintf("%s\n", respStr)
}
