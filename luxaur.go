package main

import (
	"fmt"
	"os"

	"github.com/linklux/luxaur/command"
	"github.com/mgutz/ansi"
)

func printError(err string) {
	fmt.Println(ansi.Color(err, "red"))
}

var commands = map[string]command.ICommand{
	"find":   command.NewFindCommand(),
	"search": command.NewSearchCommand(),
}

func printUsage() {
	fmt.Println("Printing usage")
}

func main() {
	args := os.Args[1:]

	command := ""
	if len(args) > 0 {
		if _, ok := commands[args[0]]; ok {
			command = args[0]
		} else {
			printError(fmt.Sprintf("Command '%s' is not supported\n", args[0]))
		}
	}

	commandArgs := []string{}
	if len(args) > 1 {
		commandArgs = args[1:]
	}

	// Try to parse command flags for the given command. Will terminate program
	// execution and print usage for the given command when an error occures.
	commands[command].ParseFlags(args[2:])
	commands[command].Execute(commandArgs)
}
