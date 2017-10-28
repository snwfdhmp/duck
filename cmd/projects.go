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
	"path/filepath"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	Yellow  func(...interface{}) string = color.New(color.FgYellow).Sprint
	Blue    func(...interface{}) string = color.New(color.FgCyan).Sprint
	Green   func(...interface{}) string = color.New(color.FgGreen).Sprint
	Red     func(...interface{}) string = color.New(color.FgRed).Sprint
	Magenta func(...interface{}) string = color.New(color.FgMagenta).Sprint
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := GetDataProjects()
		if err != nil {
			color.Red("Could not access/read duck data :" + err.Error())
			return
		}

		for _, p := range projects {
			fmt.Println("-", Blue(p.Name()), "=>", Yellow(p.Value()))
		}

		fmt.Println(len(projects), "projects.")
	},
}

func GetDataProjects() ([]*ini.Key, error) {
	config, err := getDuckData()
	if err != nil {
		return nil, err
	}

	projectsSection, err := config.GetSection("projects")
	if err != nil {
		return nil, err
	}

	projects := projectsSection.Keys()
	return projects, nil
}

func IsHealthy(path string) (bool, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}
	exists, err := afero.Exists(fs, absolutePath+"/.duck")
	if err != nil {
		return false, err
	} else if !exists {
		return false, nil
	}

	exists, err = afero.Exists(fs, absolutePath+"/.duck/conf.ini")
	if err != nil {
		return false, err
	} else if !exists {
		return false, nil
	}

	exists, err = afero.Exists(fs, absolutePath+"/.duck/packages")
	if err != nil {
		return false, err
	} else if !exists {
		return false, nil
	}

	config, err := ini.Load(absolutePath + "/.duck/conf.ini")
	if err != nil {
		return false, nil
	}

	projectSection, err := config.GetSection("project")
	if err != nil {
		return false, nil
	}

	_, err = projectSection.GetKey("name")
	if err != nil {
		return false, nil
	}

	pathKey, err := projectSection.GetKey("path")
	if err != nil {
		return false, nil
	}
	if pathKey.Value() != absolutePath {
		return false, nil
	}

	return true, nil
}

func RepairProject(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return
	}
	InitProject(initCmd, []string{}, absolutePath, false, true)
}

func init() {
	RootCmd.AddCommand(projectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
