package command

import (
	"fmt"

	"github.com/linklux/luxaur/aur_util"
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
	c.commandUtil.parseFlags("search", args, c.flags)
}

func (c *SearchCommand) Execute(args []string) bool {
	if len(args) == 0 {
		c.printError("Package search requires an argument")
		c.PrintUsage()
		return false
	}

	count, packages := aur_util.Search(args[0])

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
	c.printUsage(c.GetDescription(), c.flags)
}

func (c *SearchCommand) GetDescription() string {
	return "Search the AUR for packages that match the given name fully or partially"
}
