package interactive

import (
	"github.com/altay13/godatastructures/queue"
	"github.com/altay13/vertera/ui"
)

type HISTORY struct {
	ui           *ui.InteractiveUI
	archivedCmds *queue.Queue
}

func HistoryCommand(ui *ui.InteractiveUI, cmds *queue.Queue) InterCMD {
	return &HISTORY{
		ui:           ui,
		archivedCmds: cmds,
	}
}

func (cmd *HISTORY) Validate() error {
	return nil
}

func (cmd *HISTORY) Run() {
	if err := cmd.Validate(); err != nil {
		cmd.ui.Error(err.Error())
		return
	}

	for val := cmd.archivedCmds.Dequeue(); val != nil; val = cmd.archivedCmds.Dequeue() {
		cmd.ui.Output(val.(string))
	}

}
