package interactive

type HELP struct {
	args []string
}

func HelpCommand(args []string) InterCMD {
	return &HELP{
		args: args,
	}
}

func (cmd *HELP) Validate() error {
	return nil
}

func (cmd *HELP) Run() string {

	if err := cmd.Validate(); err != nil {
		return err.Error()
	}

	return "help command is performed"
}
