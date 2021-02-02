package command

import (
	"fmt"

	"github.com/linklux/luxaur/http_client"
)

type InfoCommand struct {
	*commandUtil

	flags map[string]*commandFlag
}

func NewInfoCommand() *InfoCommand {
	flags := map[string]*commandFlag{}

	return &InfoCommand{&commandUtil{}, flags}
}

func (c *InfoCommand) ParseFlags(args []string) {
	c.parseFlags(args, c.flags)
}

func (c *InfoCommand) Execute(args []string) bool {
	if len(args) == 0 {
		c.printError("Package info requires an argument")
		c.PrintUsage()
		return false
	}

	// TODO Support RPC's multiinfo feature.
	client := http_client.AurClient{}
	count, packages := client.Find(args)

	if count == 0 {
		c.printError(fmt.Sprintf("No package(s) found for '%v'", args))
		return false
	}

	for _, element := range packages {
		fmt.Println(element)
	}

	return true
}

func (c *InfoCommand) PrintUsage() {
	c.printUsage(c.GetDescription(), c.flags)
}

func (c *InfoCommand) GetDescription() string {
	return "Search the AUR to find a package that is an exact match with the given name"
}
