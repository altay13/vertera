package cli

import (
	"fmt"
	"log"
	"os"
)

// CLI ...
type CLI interface {
	Validate() bool
	Run() error
}

// Parse ...
func Parse() {
	if len(os.Args) < 2 {
		log.Printf("Please provide a command [ test ]")
		os.Exit(1)
	}

	var coreCLI CLI

	switch os.Args[1] {
	case "interactive":
		coreCLI = InteractiveCommand()
	case "test":
		coreCLI = TestCommand()
	default:
		os.Exit(1)
	}

	err := coreCLI.Run()
	if err != nil {
		fmt.Println(err)
	}
}
