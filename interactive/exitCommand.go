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

func (cmd *EXIT) Run() {

	os.Exit(0)

	if err := cmd.Validate(); err != nil {
		err.Error()
	}
}
