package interactive

import "os"

type EXIT struct {
}

func ExitCommand() InterCMD {
	return &EXIT{}
}

func (cmd *EXIT) Validate() error {
	return nil
}

func (cmd *EXIT) Run() string {

	os.Exit(1)

	if err := cmd.Validate(); err != nil {
		return err.Error()
	}

	return "exit command is performed"
}
