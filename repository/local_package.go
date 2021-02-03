package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/mgutz/ansi"
)

type LocalPackage struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	InstalledAt int64  `json:"installedAt"`
}

func (p LocalPackage) String() string {
	return fmt.Sprintf("%s %s %s\n  %s",
		fmt.Sprintf("%s", ansi.Color(p.Name, "blue+b")),
		fmt.Sprintf("%s", ansi.Color(p.Version, "green+b")),
		fmt.Sprintf("%s", ansi.Color(p.Description, "white")),
		fmt.Sprintf("%s", ansi.Color(fmt.Sprintf("Installed: %s", time.Unix(p.InstalledAt, 0).Format("2006-01-02")), "white")),
	)
}

// TODO Handle reading/writing in a more efficient way
type LocalPackageRepository struct {
	packages map[string]LocalPackage
}

func NewLocalPackageRepository() (*LocalPackageRepository, error) {
	base, _ := os.UserHomeDir()
	path := fmt.Sprintf("%s/.luxaur/packages.json", base)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Package list not found, (re)creating")
		err := ioutil.WriteFile(path, []byte("[]"), 644)

		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	var packages []LocalPackage

	if err := json.Unmarshal(data, &packages); err != nil {
		return nil, err
	}

	pkgMap := make(map[string]LocalPackage, len(packages))
	for _, pkg := range packages {
		pkgMap[pkg.Name] = pkg
	}

	return &LocalPackageRepository{pkgMap}, nil
}

func (repo *LocalPackageRepository) All() map[string]LocalPackage {
	return repo.packages
}

func (repo *LocalPackageRepository) Save(pkg *LocalPackage) {
	repo.packages[pkg.Name] = *pkg
}

func (repo *LocalPackageRepository) Flush() error {
	// When set to higher than zero, an empty LocalPackage is added
	s := make([]LocalPackage, 0)
	for _, pkg := range repo.packages {
		s = append(s, pkg)
	}

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	base, _ := os.UserHomeDir()
	path := fmt.Sprintf("%s/.luxaur/packages.json", base)

	err = ioutil.WriteFile(path, data, 644)

	return err
}
