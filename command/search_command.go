package command

import (
	"fmt"

	"github.com/linklux/luxaur/http_client"
)

type SearchCommand struct {
	*commandUtil

	flags map[string]*commandFlag
}

func NewSearchCommand() *SearchCommand {
	flags := map[string]*commandFlag{
		"no-deprecation": &commandFlag{"no-deprecation", "d", "bool", "Do not show deprecated packages", false},
	}

	return &SearchCommand{&commandUtil{}, flags}
}

func (c *SearchCommand) ParseFlags(args []string) {
	c.commandUtil.parseFlags(args, c.flags)
}

func (c *SearchCommand) Execute(args []string) bool {
	for _, element := range c.flags {
		fmt.Println(element)
	}

	if len(args) == 0 {
		c.printError("Find command requires an argument")
		return false
	}

	client := http_client.AurClient{}
	count, packages := client.Search(args[0])

	if count == 0 {
		c.printError(fmt.Sprintf("No packages found for '%s'", args[0]))
		return false
	}

	noDeprecation := c.flags["no-deprecation"].Value.(bool)
	for _, element := range packages {
		if noDeprecation && element.Outdated > 0 {
			continue
		}

		fmt.Println(element)
	}

	return true
}

func (c *SearchCommand) PrintUsage() {
	c.printUsage(c.flags)
}
