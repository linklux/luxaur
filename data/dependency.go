package data

import (
	"fmt"
)

type Dependency struct {
	Name    string
	Version string
}

func (d Dependency) String() string {
	return fmt.Sprintf("%s %s", d.Name, d.Version)
}
