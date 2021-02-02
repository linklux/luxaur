package data

import (
	"fmt"
	"time"

	"github.com/mgutz/ansi"
)

type AurPackageSearch struct {
	Id          int    `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Version     string `json:"Version"`
	Outdated    int64  `json:"OutOfDate"`
}

func (p AurPackageSearch) String() string {
	outdatedStr := ""

	if p.Outdated != 0 {
		date := time.Unix(p.Outdated, 0).Format("2006-01-02")
		outdatedStr = fmt.Sprintf("%s", "(Outdated as of "+date+")")
	}

	return fmt.Sprintf("%s %s %s\n    %s",
		fmt.Sprintf("%s", ansi.Color(p.Name, "blue+b")),
		fmt.Sprintf("%s", ansi.Color(p.Version, "green+b")),
		fmt.Sprintf("%s", ansi.Color(outdatedStr, "red+b")),
		fmt.Sprintf("%s", ansi.Color(p.Description, "white")),
	)
}
