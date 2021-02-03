package command

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/linklux/luxaur/http_client"
	"github.com/linklux/luxaur/io_util"
	"github.com/linklux/luxaur/model"
	"github.com/linklux/luxaur/repository"
)

type InstallCommand struct {
	*commandUtil

	flags map[string]*commandFlag
}

func NewInstallCommand() *InstallCommand {
	flags := map[string]*commandFlag{
		"skip-pgp": &commandFlag{"skip-pgp", "s", "bool", "Skip PGP signature checks", false},
	}

	return &InstallCommand{&commandUtil{}, flags}
}

func (c *InstallCommand) ParseFlags(args []string) {
	c.parseFlags("install", args, c.flags)
}

func (c *InstallCommand) Execute(args []string) bool {
	if len(args) == 0 {
		c.printError("Package install requires an argument")
		c.PrintUsage()
		return false
	}

	client := http_client.AurClient{}
	count, packages := client.Find(args)

	if count == 0 {
		c.printError(fmt.Sprintf("No package(s) found for '%v'", args))
		return false
	}

	for _, pkg := range packages {
		c.install(&pkg)
	}

	return true
}

// TODO Relocate to reuse for updating packages (composited?)
func (c *InstallCommand) install(pkg *model.AurPackageInfo) {
	file, err := http_client.Download(pkg)
	if err != nil {
		c.printError(fmt.Sprintf("Failed to download package for '%s': %s", pkg.Name, err.Error()))
		return
	}

	err = io_util.Untargz(file)
	if err != nil {
		c.printError(fmt.Sprintf("Failed to extract package for '%s': '%s'", pkg.Name, err.Error()))
		return
	}

	executable, _ := exec.LookPath("makepkg")
	args := []string{executable, "-si"}

	if c.flags["skip-pgp"].Value.(bool) {
		args = append(args, "--skippgpcheck")
	}

	cmd := &exec.Cmd{
		Path:   executable,
		Args:   args,
		Dir:    path.Dir(file) + "/" + pkg.Name,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
	}

	if err := cmd.Run(); err != nil {
		c.printError(err.Error())
		return
	}

	// TODO Check if current version already installed, prompt for install confirmation when it is
	repo, err := repository.NewLocalPackageRepository()
	if err != nil {
		c.printError(err.Error())
		return
	}

	repo.Save(&repository.LocalPackage{pkg.Name, pkg.Version, pkg.Description, time.Now().Unix()})
	err = repo.Flush()

	if err != nil {
		c.printError(err.Error())
	}
}

func (c *InstallCommand) PrintUsage() {
	c.printUsage(c.GetDescription(), c.flags)
}

func (c *InstallCommand) GetDescription() string {
	return "Install one or more packages from the AUR"
}
