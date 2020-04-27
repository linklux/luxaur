package command

import (
	"flag"
	"fmt"

	"github.com/mgutz/ansi"
)

type ICommand interface {
	ParseFlags(args []string)
	Execute(args []string) bool

	PrintUsage()
}

type commandFlag struct {
	Name        string
	Shortname   string
	Datatype    string
	Description string
	Value       interface{}
}

type commandUtil struct{}

func (c *commandUtil) parseFlags(args []string, flags map[string]*commandFlag) {
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	fs.Usage = func() {
		c.printUsage(flags)
	}

	// TODO There must be a better way to do this...
	bools := map[string]*bool{}
	ints := map[string]*int{}
	strings := map[string]*string{}

	for key, element := range flags {
		switch element.Datatype {
		case "bool":
			if defaultValue, ok := element.Value.(bool); ok {
				bools[key] = &defaultValue

				if element.Shortname != "" {
					fs.BoolVar(bools[key], element.Shortname, defaultValue, element.Description)
				}

				fs.BoolVar(bools[key], element.Name, defaultValue, element.Description)
			}

		case "int":
			if defaultValue, ok := element.Value.(int); ok {
				ints[key] = &defaultValue

				if element.Shortname != "" {
					fs.IntVar(ints[key], element.Shortname, defaultValue, element.Description)
				}

				fs.IntVar(ints[key], element.Name, defaultValue, element.Description)
			}

		case "string":
			if defaultValue, ok := element.Value.(string); ok {
				strings[key] = &defaultValue

				if element.Shortname != "" {
					fs.StringVar(strings[key], element.Shortname, defaultValue, element.Description)
				}

				fs.StringVar(strings[key], element.Name, defaultValue, element.Description)
			}

		}
	}

	fs.Parse(args)

	for key, element := range bools {
		flags[key].Value = *element
	}

	for key, element := range ints {
		flags[key].Value = *element
	}

	for key, element := range strings {
		flags[key].Value = *element
	}
}

func (c *commandUtil) printUsage(flags map[string]*commandFlag) {
	fmt.Println("TODO!")
}

func (c *commandUtil) printError(err string) {
	fmt.Println(ansi.Color(err, "red"))
}
