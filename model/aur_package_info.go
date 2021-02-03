package model

import (
	"fmt"
	"time"

	"github.com/mgutz/ansi"
)

type AurPackageInfo struct {
	AurPackageSearch

	FirstSubmitted       int64    `json:"FirstSubmitted"`
	LastModified         int64    `json:"LastModified"`
	Url                  string   `json:"URLPath"`
	Dependencies         []string `json:"Depends"`
	OptionalDependencies []string `json:"OptDepends"`
	Keywords             []string `json:"Keywords"`
}

func (p AurPackageInfo) String() string {
	outdatedStr := ""
	if p.Outdated != 0 {
		date := time.Unix(p.Outdated, 0).Format("2006-01-02")
		outdatedStr = fmt.Sprintf("%s", "(Outdated as of "+date+")")
	}

	dependencies := ""
	for _, element := range p.Dependencies {
		dependencies += fmt.Sprintf("    - %s\n", ansi.Color(element, "blue"))
	}

	optional := ""
	for _, element := range p.OptionalDependencies {
		optional += fmt.Sprintf("    - %s\n", ansi.Color(element, "blue"))
	}

	return fmt.Sprintf("%s %s %s\n  %s\n\n  %s\n  %s",
		fmt.Sprintf("%s", ansi.Color(p.Name, "blue+b")),
		fmt.Sprintf("%s", ansi.Color(p.Version, "green+b")),
		fmt.Sprintf("%s", ansi.Color(outdatedStr, "red+b")),
		fmt.Sprintf("%s", ansi.Color(p.Description, "white")),
		fmt.Sprintf("%s\n%s", ansi.Color("Dependencies", "white+b"), ansi.Color(dependencies, "white")),
		fmt.Sprintf("%s\n%s", ansi.Color("Optional dependencies", "white+b"), ansi.Color(optional, "white")),
	)
}
