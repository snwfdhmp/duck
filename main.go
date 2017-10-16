// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/snwfdhmp/duck/cmd"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	fs            afero.Fs
	packagesPaths []string
	duckCommands  []duckCommand
)

const packagesExtension = ".duckpkg.ini"

func main() {
	fs = afero.NewOsFs()
	exists, err := afero.Exists(fs, ".duck/conf.ini")
	if err != nil {
		color.Red(err.Error())
		return
	}
	if !exists {
		//color.Yellow("Warning: not a duck repository")
		cmd.Execute()
		return
	}

	cfg, err := ini.Load(".duck/conf.ini")
	if err != nil {
		color.Red(err.Error())
		return
	}
	cfgPackages, err := cfg.GetSection("packages")
	if err != nil {
		color.Red(err.Error())
		return
	}
	includePath, err := cfgPackages.GetKey("directory")
	if err != nil {
		color.Red(err.Error())
		return
	}

	packagesPaths = getPackagesPaths(".duck/" + includePath.String())
	duckCommands, err = scanCommands(packagesPaths)
	if err != nil {
		color.Red(err.Error())
	}

	createCobraCommands(duckCommands)

	cmd.Execute()
}

func getPackagesPaths(includePath string) []string {
	var paths []string
	//color.Yellow("Including from: " + includePath)
	fi, err := afero.ReadDir(fs, includePath)
	if err != nil {
		color.Red(err.Error())
		return []string{}
	}
	var path string
	for i := 0; i < len(fi); i++ {
		path = includePath + "/" + fi[i].Name()
		if fi[i].IsDir() {
			paths = append(paths, getPackagesPaths(path)...)
			continue
		}
		if path[len(path)-len(packagesExtension):] == packagesExtension {
			paths = append(paths, path)
		}
	}
	return paths
}

type duckCommand struct {
	Name      string
	Shortcut  string
	Cmd       string
	ShortHelp string
	LongHelp  string
}

func scanCommands(paths []string) ([]duckCommand, error) {
	var cmds []duckCommand
	for i := 0; i < len(paths); i++ {
		cfg, err := ini.Load(paths[i])
		if err != nil {
			color.Red(err.Error())
		}
		sections := cfg.Sections()
		for j := 0; j < len(sections); j++ {
			if sections[j].Name() == "DEFAULT" {
				continue
			}
			cmds = append(cmds, duckCommand{
				Name:      sections[j].Name(),
				Cmd:       sections[j].Key("cmd").String(),
				Shortcut:  sections[j].Key("shortcut").String(),
				ShortHelp: sections[j].Key("help").String(),
				LongHelp:  sections[j].Key("longHelp").String(),
			})
		}
	}
	return cmds, nil
}

func createCobraCommands(cmds []duckCommand) {
	for i := 0; i < len(cmds); i++ {
		var tmpCmd = &cobra.Command{
			Use:     cmds[i].Name,
			Short:   cmds[i].ShortHelp,
			Long:    cmds[i].LongHelp,
			Aliases: []string{cmds[i].Shortcut},
			Run: func(cmd *cobra.Command, args []string) {
				i := cmd.DuckCmdIndex
				execCmd := exec.Command("sh", "-c", cmds[i].Cmd)
				execCmd.Stderr = os.Stderr
				execCmd.Stdout = os.Stdout
				execCmd.Stdin = os.Stdin
				err := execCmd.Run()
				if err != nil {
					color.Red("An error occured while executing this command. Error: " + err.Error())
					os.Exit(1)
				}
			},
			DuckCmdIndex: i,
		}
		cmd.RootCmd.AddCommand(tmpCmd)
	}
}
