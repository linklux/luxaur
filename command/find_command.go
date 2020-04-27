package command

import (
	"fmt"

	"github.com/linklux/luxaur/http_client"
)

type FindCommand struct {
	*commandUtil

	flags map[string]*commandFlag
}

func NewFindCommand() *FindCommand {
	flags := map[string]*commandFlag{}

	return &FindCommand{&commandUtil{}, flags}
}

func (c *FindCommand) ParseFlags(args []string) {
	c.parseFlags(args, c.flags)
}

func (c *FindCommand) Execute(args []string) bool {
	if len(args) == 0 {
		c.printError("Find command requires an argument")
		return false
	}

	client := http_client.AurClient{}
	count, pkg := client.Find(args[0])

	if count == 0 {
		c.printError(fmt.Sprintf("No package found for '%s'", args[0]))
		return false
	}

	fmt.Println(pkg)
	return true
}

func (c *FindCommand) PrintUsage() {
	c.printUsage(c.flags)
}
