package interactive

import (
	"fmt"
	"strings"

	"github.com/altay13/godatastructures/queue"
	"github.com/altay13/vertera/eventHandler"
	"github.com/altay13/vertera/eventHandler/hazelcast"
	"github.com/altay13/vertera/eventHandler/redis"
	"github.com/altay13/vertera/eventHandler/tarantool"
	"github.com/altay13/vertera/ui"
)

type Interactive struct {
	ui *ui.InteractiveUI

	handler *eventHandler.EventHandler

	cmdCh chan string

	archivedCmds *queue.Queue
}

type InterCMD interface {
	Validate() error
	Run()
}

func NewInteractive() *Interactive {
	inter := &Interactive{
		ui:           ui.DefaultInteractiveUI(),
		cmdCh:        make(chan string, 1),
		archivedCmds: queue.NewQueue(),
	}

	inter.archivedCmds.SetSize(10)

	return inter
}

func (inter *Interactive) Start() {
	go func() {
		inter.ui.StartInteractive(inter.cmdCh)
	}()

	inter.ui.Info("Interactive mode console:")

	for {
		select {
		case cmd := <-inter.cmdCh:
			inter.archivedCmds.Enqueue(cmd)
			inter.parseCMD(cmd)
		}
	}
}

func (inter *Interactive) SetDatabase(dbName string, configStr string) {
	var db eventHandler.EventStore
	inter.handler.CloseDB()
	var err error
	switch dbName {
	case eventHandler.REDIS:
		conf := redis.DefaultConfig()
		inter.askConfig("(example: Host=localhost:6379;PoolIdleSize=5;IdleTimeout=120)", conf, configStr)
		inter.ui.Info(fmt.Sprintf("Configuration set: Host=%s; PoolIdleSize=%d; IdleTimeout=%d", conf.Host, conf.IdleTimeout, conf.PoolIdleSize))
		db, err = redis.NewRedis(conf)
	case eventHandler.CASSANDRA:
	case eventHandler.ROCKSDB:
	case eventHandler.HAZELCAST:
		conf := hazelcast.DefaultConfig()
		inter.askConfig("(example: Host=localhost:5701)", conf, configStr)
		inter.ui.Info(fmt.Sprintf("Configuration set: Host=%s", conf.Host))
		db, err = hazelcast.NewHazelcast(conf)
	case eventHandler.TARANTOOL:
		conf := tarantool.DefaultConfig()
		inter.askConfig("(example: Host=localhost:3301;Timeout=500;Reconnect=1;MaxReconnects=3;User=test;Pass=test;Space=tester)", conf, configStr)
		inter.ui.Info(fmt.Sprintf("Configuration set: Host=%s; Timeout=%d; Reconnect=%d; MaxReconnects=%d; User=%s; Pass=%s; Space=%s",
			conf.Host, conf.Timeout, conf.Reconnect, conf.MaxReconnects, conf.User, conf.Pass, conf.Space))
		db, err = tarantool.NewTarantool(conf)
	default:
		inter.ui.Error(fmt.Sprintf("There is no such database: %s", dbName))
		return
	}
	if err != nil {
		inter.ui.Error(fmt.Sprintf("Failed to connect to db: %s. Err - %s", dbName, err))
		return
	}
	inter.ui.Info(fmt.Sprintf("Set to database: %s", dbName))
	inter.handler = eventHandler.NewEventHandler(db)
}

func (inter *Interactive) askConfig(dbStr string, conf eventHandler.EventStoreConfig, configStr string) {
	str := fmt.Sprintf("Please provide config string or leave empty if you want to use default settings: %s", dbStr)
	if len(configStr) <= 0 {
		configStr, _ = inter.ui.Ask(str)
	}
	for err := conf.SetByConfigString(configStr); err != nil; configStr, _ = inter.ui.Ask(str) {
		inter.ui.Error(err.Error())
		err = conf.SetByConfigString(configStr)
	}
}

func (inter *Interactive) parseCMD(cmd string) {
	// format the command.
	cmds := strings.Fields(strings.Replace(cmd, " = ", "=", -1))
	if len(cmds) <= 0 {
		inter.ui.Info("Please enter the command. Type `help` if you don't know what to do.")
		return
	}

	var coreInterCmd InterCMD

	switch Command(cmds[0]) {
	case cmdUSE:
		// use is for changing the backend DB [redis, cassandra, rocksdb]
		// TODO: think about configuration. Right now temp solution with default configuration
		inter.SetDatabase(cmds[1], strings.Join(cmds[2:], " "))
		return
	case cmdVAR:
		// var is for variable set. I have to persist the var in interactive session
	case cmdSET:
		coreInterCmd = SetCommand(cmds[1:], inter.handler, inter.ui)
	case cmdGET:
		coreInterCmd = GetCommand(cmds[1:], inter.handler, inter.ui)
	case cmdHISTORY:
		coreInterCmd = HistoryCommand(inter.ui, inter.archivedCmds)
	case cmdHELP:
		coreInterCmd = HelpCommand(cmds[1:], inter.ui)
	case cmdEXIT:
		coreInterCmd = ExitCommand()
	default:
		inter.ui.Error("Unknown command. Type `help` if you don't know what to do.")
		return
	}

	coreInterCmd.Run()
}
