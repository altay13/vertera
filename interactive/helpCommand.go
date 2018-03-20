package interactive

import (
	"fmt"

	"github.com/altay13/vertera/ui"
)

type HELP struct {
	args []string
	ui   *ui.InteractiveUI
}

func HelpCommand(args []string, ui *ui.InteractiveUI) InterCMD {
	return &HELP{
		args: args,
		ui:   ui,
	}
}

func (cmd *HELP) Validate() error {
	if len(cmd.args) == 1 {
		if len(CMDs[Command(cmd.args[0])]) > 0 {
			return nil
		}
	} else if len(cmd.args) == 0 {
		return nil
	}

	return fmt.Errorf("Failed to parse command. %s", CMDs[cmdHELP])
}

func (cmd *HELP) Run() {
	if err := cmd.Validate(); err != nil {
		cmd.ui.Error(err.Error())
		return
	}

	if len(cmd.args) == 1 {
		cmd.ui.Info(CMDs[Command(cmd.args[0])])
		return
	}

	for _, v := range CMDs {
		cmd.ui.Info(v)
	}

}
