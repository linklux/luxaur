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
	"list":    command.NewListCommand(),
	"info":    command.NewInfoCommand(),
	"search":  command.NewSearchCommand(),
	"install": command.NewInstallCommand(),
}

func printUsage() {
	fmt.Printf("%s\n%s\n\n",
		"A simple lightweight AUR tool that makes searching, installing and managing installed packages easier.",
		"Usage: luxaur [command] <arguments> <flags>",
	)

	fmt.Println("Available commands:")
	for name, element := range commands {
		fmt.Printf("%s\t%s\n", name, element.GetDescription())
	}

	fmt.Println("\nFor detailed usage of a command, use: luxaur [command] -h|--help")
}

func main() {
	args := os.Args[1:]

	command := ""
	if len(args) > 0 {
		if _, ok := commands[args[0]]; ok {
			command = args[0]
		} else {
			printError(fmt.Sprintf("Command '%s' is not supported\n", args[0]))
			printUsage()
			return
		}
	} else {
		printUsage()
		return
	}

	commandArgs := []string{}
	if len(args) > 1 {
		commandArgs = args[1:]

		if commandArgs[0] == "-h" || commandArgs[0] == "--help" {
			commands[command].PrintUsage()
			return
		}
	}

	commandFlags := []string{}
	if len(args) > 2 {
		commandFlags = args[2:]
	}

	// TODO Allow multiple arguments per command where desired.
	// Try to parse command flags for the given command. Will terminate program
	// execution and print usage for the given command when an error occures.
	commands[command].ParseFlags(commandFlags)
	commands[command].Execute(commandArgs)
}
