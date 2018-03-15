package cli

import (
	"flag"
	"log"
	"os"

	"github.com/altay13/vertera/interactive"
)

// InteractiveCLI ...
type InteractiveCLI struct {
	ArgsPtr map[string]*string

	dbs        map[string]bool
	selectedDB string
	config     string
}

// InteractiveCommand ...
func InteractiveCommand() CLI {
	interactiveCommand := flag.NewFlagSet("interactive", flag.ExitOnError)

	interactiveCLI := &InteractiveCLI{
		dbs: map[string]bool{"redis": true, "cassandra": true, "rocksdb": true},
	}
	interactiveCLI.ArgsPtr = make(map[string]*string)
	interactiveCLI.ArgsPtr["db"] = interactiveCommand.String("db", "", "Specify the key/value database.(Optional)")
	interactiveCLI.ArgsPtr["config"] = interactiveCommand.String("config", "", "Specify the connection string to database.(Optional)")

	interactiveCommand.Parse(os.Args[2:])

	if interactiveCommand.Parsed() {
		if !interactiveCLI.Validate() {
			interactiveCommand.PrintDefaults()
			os.Exit(1)
		}
	}

	return interactiveCLI
}

// Validate ...
func (p *InteractiveCLI) Validate() bool {
	if len(*p.ArgsPtr["db"]) > 0 {
		if !p.dbs[*p.ArgsPtr["db"]] {
			log.Printf("[ERR] - Failed to initiate %s interactive console.\n", *p.ArgsPtr["db"])
			return false
		}
		if len(*p.ArgsPtr["config"]) <= 0 {
			log.Printf("[ERR] - Please provide the connection configuration string(--config='').\n")
			return false
		}

		p.selectedDB = *p.ArgsPtr["db"]
	}

	return true
}

// Run ...
func (p *InteractiveCLI) Run() error {
	inter := interactive.NewInteractive()
	if len(p.selectedDB) > 0 {
		// TODO: add set database method
		// inter.SetDatabase(dbname, config)
	}
	inter.Start()
	return nil
}
