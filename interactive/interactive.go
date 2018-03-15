package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/altay13/vertera/eventHandler"
	"github.com/altay13/vertera/eventHandler/redis"
)

type Interactive struct {
	stopChan chan bool
	reader   *bufio.Reader

	handler *eventHandler.EventHandler
}

type InterCMD interface {
	Validate() error
	Run() string
}

func NewInteractive() *Interactive {
	inter := &Interactive{
		stopChan: make(chan bool, 1),
		reader:   bufio.NewReader(os.Stdin),
	}

	// TODO: Just for test hardcode the redis db
	conf := redis.DefaultConfig()
	db := redis.NewRedis(conf)

	inter.handler = eventHandler.NewEventHandler(db)

	return inter
}

func (inter *Interactive) Start() {
	scanner := bufio.NewScanner(inter.reader)

	for scanner.Scan() {
		cmd := scanner.Text()
		inter.parseCMD(cmd)
	}
}

func (inter *Interactive) SetDatabase(dbName string, config string) {
}

func (inter *Interactive) parseCMD(cmd string) {
	// format the command.
	cmds := strings.Fields(strings.Replace(cmd, " = ", "=", -1))
	if len(cmds) <= 0 {
		fmt.Println("Please enter the command. Type help if you don't know what to do.")
	}

	var coreInterCmd InterCMD

	switch cmds[0] {
	case "use":
		// use is for changing the backend DB [redis, cassandra, rocksdb]
		// db := UseCommand(cmds[1:])
		return
	case "var":
		// var is for variable set. I have to persist the var in interactive session
	case "set":
		coreInterCmd = SetCommand(cmds[1:], inter.handler)
	case "get":
		coreInterCmd = GetCommand(cmds[1:], inter.handler)
	case "help":
		coreInterCmd = HelpCommand(cmds)
	case "exit":
		coreInterCmd = ExitCommand()
	default:
		fmt.Println("Unknown command. Type help if you don't know what to do.")
		return
	}

	output := coreInterCmd.Run()

	fmt.Println(output)
}
