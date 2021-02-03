package command

import (
	"fmt"

	"github.com/linklux/luxaur/repository"
)

type ListCommand struct {
	*commandUtil

	flags map[string]*commandFlag
}

func NewListCommand() *ListCommand {
	flags := map[string]*commandFlag{}

	return &ListCommand{&commandUtil{}, flags}
}

func (c *ListCommand) ParseFlags(args []string) {
	c.parseFlags("list", args, c.flags)
}

func (c *ListCommand) Execute(args []string) bool {
	repo, err := repository.NewLocalPackageRepository()
	if err != nil {
		c.printError(err.Error())
		return false
	}

	for _, element := range repo.All() {
		fmt.Println(element)
	}

	return true
}

func (c *ListCommand) PrintUsage() {
	c.printUsage(c.GetDescription(), c.flags)
}

func (c *ListCommand) GetDescription() string {
	return "List all installed AUR packages known by luxaur"
}
