package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/altay13/vertera/eventHandler"
	"github.com/altay13/vertera/eventHandler/hazelcast"
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
	var db eventHandler.EventStore
	// TODO: Just for test hardcode the redis db
	switch dbName {
	case eventHandler.REDIS:
		if inter.handler != nil {
			if inter.handler.GetDBName() != eventHandler.REDIS {
				inter.handler.CloseDB()
			}
		}
		db = redis.NewRedis(redis.DefaultConfig())
	case eventHandler.CASSANDRA:
	case eventHandler.ROCKSDB:
	case eventHandler.HAZELCAST:
		if inter.handler != nil {
			if inter.handler.GetDBName() != eventHandler.REDIS {
				inter.handler.CloseDB()
			}
		}
		db = hazelcast.NewHazelcast(hazelcast.DefaultConfig())
	default:
		fmt.Printf("There is no such database: %s\n", dbName)
		return
	}
	fmt.Printf("Set to database: %s\n", dbName)
	inter.handler = eventHandler.NewEventHandler(db)
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
		// TODO: think about configuration. Right now temp solution with default configuration
		inter.SetDatabase(cmds[1], strings.Join(cmds[1:], " "))
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
