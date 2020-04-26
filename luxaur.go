package main

import (
	"fmt"
	"os"

	"github.com/linklux/luxaur/http_client"
	"github.com/mgutz/ansi"
)

// TODO Automated validation for commands.
type Command interface {
	execute(args []string)
}

type UsageCommand struct{}
type FindCommand struct{}
type SearchCommand struct{}

func errorOutput(err string) {
	fmt.Println(ansi.Color(err, "red"))
}

func (c UsageCommand) execute(args []string) {
	fmt.Println("Running usage")
}

func (c FindCommand) execute(args []string) {
	if len(args) == 0 {
		errorOutput("Find command requires an argument")
		return
	}

	client := http_client.AurClient{}
	count, pkg := client.Find(args[0])

	if count == 0 {
		errorOutput(fmt.Sprintf("No package found for '%s'", args[0]))
		return
	}

	fmt.Println(pkg)
}

func (c SearchCommand) execute(args []string) {
	if len(args) == 0 {
		errorOutput("Find command requires an argument")
		return
	}

	client := http_client.AurClient{}
	count, packages := client.Search(args[0])

	if count == 0 {
		errorOutput(fmt.Sprintf("No packages found for '%s'", args[0]))
		return
	}

	for _, element := range packages {
		fmt.Println(element)
	}
}

var commands = map[string]Command{
	"":       &UsageCommand{},
	"find":   &FindCommand{},
	"search": &SearchCommand{},
}

func main() {
	args := os.Args[1:]

	command := ""
	if len(args) > 0 {
		if _, ok := commands[args[0]]; ok {
			command = args[0]
		} else {
			errorOutput(fmt.Sprintf("Command '%s' is not supported\n", args[0]))
		}
	}

	commandArgs := []string{}
	if len(args) > 1 {
		commandArgs = args[1:]
	}

	commands[command].execute(commandArgs)
}
