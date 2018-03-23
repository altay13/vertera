package interactive

type Command string

const (
	cmdUSE     Command = "use"
	cmdVAR     Command = "var"
	cmdSET     Command = "set"
	cmdGET     Command = "get"
	cmdHISTORY Command = "history"
	cmdHELP    Command = "help"
	cmdEXIT    Command = "exit"
)

var (
	CMDs map[Command]string = map[Command]string{
		cmdUSE:     "use - Usage: use [redis|rocksdb|cassandra|tarantool|hazelcast].",
		cmdVAR:     "var - Sets a variable. Usage: var test = 14 (Not implemented yet).",
		cmdSET:     "set - Sets value to a key. Usage: set [key] = [value].",
		cmdGET:     "get - Gets a value by key. Usage: get [key].",
		cmdHISTORY: "history - Gets last 10 command history. Usage: history.",
		cmdHELP:    "help - Type help to get commands list. Usage: help [command].",
		cmdEXIT:    "exit - Exits interactive console.",
	}
)
