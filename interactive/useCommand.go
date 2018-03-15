package interactive

type USE struct {
	args []string

	dbName string
	config string
}

func UseCommand(args []string) InterCMD {
	return &USE{
		args: args,
	}
}

func (cmd *USE) Validate() error {
	return nil
}

func (cmd *USE) Run() string {

	if err := cmd.Validate(); err != nil {
		return err.Error()
	}

	return "use command is performed"
}
