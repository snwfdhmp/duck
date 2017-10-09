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

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	force bool

	defaultIncludePath = "packages"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()

		//if force, remove the existing .duck
		if force {
			err := fs.RemoveAll(".duck")
			if err != nil {
				color.Red(err.Error())
			}
			fmt.Println("Deleted existing repo")
		}

		//if not exists, create, else return error
		exists, err := afero.Exists(fs, ".duck")
		if err != nil {
			color.Red(err.Error())
		}
		if !exists {
			err = fs.Mkdir(".duck", 0777)
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

		wd, err := os.Getwd()
		if err != nil {
			color.Red(err.Error())
			return
		}

		wdArr := strings.Split(wd, "/")

		name, err := project.NewKey("name", wdArr[len(wdArr)-1])
		fmt.Println("Project name:", name)

		path, err := project.NewKey("path", wd)
		fmt.Println("Project path:", path)

		packages, err := cfg.NewSection("packages")
		if err != nil {
			color.Red(err.Error())
			return
		}

		includePath, err := packages.NewKey("directory", defaultIncludePath)

		fmt.Println("Creating packages directory '.duck/" + includePath.String() + "'")
		err = os.Mkdir(".duck/"+includePath.String(), 0777)

		fmt.Println("Writing configuration to .duck/conf.ini")

		err = cfg.SaveTo(".duck/conf.ini")
		if err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("Init performed successfully")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Delete .duck if it exists")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
