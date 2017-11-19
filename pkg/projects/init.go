package projects

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/snwfdhmp/duck/pkg/data"
	"github.com/spf13/afero"
)

var (
	fs = afero.NewOsFs()
)

func InitProject(args []string, path string, verbose bool, forceMode bool) {
	//if force, remove the existing .duck
	if forceMode {
		err := fs.RemoveAll(path + "/.duck")
		if err != nil {
			color.Red(err.Error())
		}
		if verbose {
			fmt.Println("Deleted existing repo")
		}
	}

	//if not exists, create, else return error
	exists, err := afero.Exists(fs, path+"/.duck")
	if err != nil {
		color.Red(err.Error())
	}
	if !exists {
		err = fs.Mkdir(path+"/.duck", 0777)
		if err != nil {
			color.Red(err.Error())
		}
	} else {
		fmt.Println("A duck repo already exists here. (-f to force)")
		return
	}

	//start config
	cfg := ini.Empty()
	project, err := cfg.NewSection("project")
	if err != nil {
		color.Red(err.Error())
		return
	}

	wdArr := strings.Split(path, "/")

	name, err := project.NewKey("name", wdArr[len(wdArr)-1])
	if verbose {
		fmt.Println("Project name:", name)
	}

	pathKey, err := project.NewKey("path", path)
	if verbose {
		fmt.Println("Project path:", pathKey)
	}

	packages, err := cfg.NewSection("packages")
	if err != nil {
		color.Red(err.Error())
		return
	}

	includePath, err := packages.NewKey("directory", PackagesDir)

	if verbose {
		fmt.Println("Creating packages directory '.duck/" + includePath.String() + "'")
	}
	err = os.Mkdir(path+"/.duck/"+includePath.String(), 0777)

	if verbose {
		fmt.Println("Writing configuration to .duck/conf.ini")
	}

	err = cfg.SaveTo(path + "/.duck/conf.ini")
	if err != nil {
		color.Red(err.Error())
		return
	}

	err = data.AddProject(name.String(), pathKey.String(), false, false)
	if err != nil {
		color.Red("Could not add project to our database. Error: " + err.Error())
	}

	if verbose {
		color.Green("Init performed successfully")
	}
}
