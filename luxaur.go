package main

import (
	"fmt"
	"os"

	"github.com/linklux/luxaur/http_client"
	"github.com/mgutz/ansi"
)

type Command interface {
	execute(args []string)
}

type UsageCommand struct{}
type SearchCommand struct{}

func (c UsageCommand) execute(args []string) {
	fmt.Println("Running usage")
}

func (c SearchCommand) execute(args []string) {
	client := http_client.AurClient{}
	packages := client.Search("spotify")

	for _, element := range packages {
		fmt.Println(element)
	}
}

var commands = map[string]Command{
	"":       &UsageCommand{},
	"search": &SearchCommand{},
}

func main() {
	args := os.Args[1:]

	command := ""
	if len(args) > 0 {
		if _, ok := commands[args[0]]; ok {
			command = args[0]
		} else {
			fmt.Println(ansi.Color(fmt.Sprintf("Command '%s' is not supported\n", args[0]), "red"))
		}
	}

	commandArgs := []string{}
	if len(args) > 1 {
		commandArgs = args[1:]
	}

	commands[command].execute(commandArgs)
}
