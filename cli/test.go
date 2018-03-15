package cli

import (
	"flag"
	"log"
	"os"
	"strings"
)

// TestCLI ...
type TestCLI struct {
	ArgsPtr map[string]*string
}

// TestCommand ...
func TestCommand() CLI {
	testCommand := flag.NewFlagSet("test", flag.ExitOnError)

	testCLI := &TestCLI{}
	testCLI.ArgsPtr = make(map[string]*string)
	testCLI.ArgsPtr["file"] = testCommand.String("file", "", "Specify file.(Required)")

	testCommand.Parse(os.Args[2:])

	if testCommand.Parsed() {
		if !testCLI.Validate() {
			testCommand.PrintDefaults()
			os.Exit(1)
		}
	}

	return testCLI
}

// Validate ...
func (p *TestCLI) Validate() bool {
	if len(*p.ArgsPtr["file"]) <= 0 {
		return false
	}

	if !strings.Contains(*p.ArgsPtr["file"], ".json") {
		log.Printf("[ERR] - Failed to locate %s file", *p.ArgsPtr["file"])
		return false
	}

	return true
}

// Run ...
func (p *TestCLI) Run() error {
	return nil
}
