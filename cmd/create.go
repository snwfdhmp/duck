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
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a ling",
	Long: `Permits you to create a new ling (duck command).
You can use it in two modes :
    - interactive: 'duck create'
    - one line: 'duck create <author> <command_name> <shortcut> <commands_to_execute> <help_message>'`,
	Run: func(cmd *cobra.Command, args []string) {
		err := loadProjectConfig()
		if err != nil {
			color.Red("Could not load project configuration : " + err.Error())
			return
		}

		var author, name, shortcut, command, help string

		if len(args) == 0 {
			author, name, shortcut, command, help, err = getArgsFromScanner()
		} else if len(args) != 5 {
			color.Red("Syntax: duck create <author> <command_name> <shortcut> <commands_to_execute> <help_message>")
			return
		} else {
			author, name, shortcut, command, help = args[0], args[1], args[2], args[3], args[4]
		}

		path, err := getConfigString("packages.directory")
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}
		configPath := ".duck/" + path + "/" + author

		fs := afero.NewOsFs()
		exists, err := afero.Exists(fs, configPath)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}
		if !exists {
			err = fs.Mkdir(configPath, 0777)
			if err != nil {
				color.Red("Error: " + err.Error())
				return
			}
		}
		configPath += "/custom.duckpkg.ini"
		exists, err = afero.Exists(fs, configPath)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}
		if !exists {
			_, err = fs.Create(configPath)
			if err != nil {
				color.Red("Error: " + err.Error())
				return
			}
		}
		config, err := ini.Load(configPath)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}
		section, err := config.NewSection(name)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}
		_, err = section.NewKey("cmd", command)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}
		_, err = section.NewKey("shortcut", shortcut)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}
		_, err = section.NewKey("help", help)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}
		err = config.SaveTo(configPath)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}

		color.Green("Ling successfully created")
	},
}

func readString(text string) (string, error) {
	fmt.Print(text)
	input, err := reader.ReadString('\n')
	return input[:len(input)-1], err
}

func init() {
	RootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getArgsFromScanner() (string, string, string, string, string, error) {
	fmt.Println("You are going to create a ling")

	author, err := readString("Enter the author name : ")
	if err != nil {
		color.Red("Error: " + err.Error())
		return "", "", "", "", "", err
	}
	name, err := readString("Choose the command name (will be used as 'duck name') : ")
	if err != nil {
		color.Red("Error: " + err.Error())
		return "", "", "", "", "", err
	}
	command, err := readString("Enter the command to execute: ")
	if err != nil {
		color.Red("Error: " + err.Error())
		return "", "", "", "", "", err
	}
	help, err := readString("Enter message to be displayed for help: ")
	if err != nil {
		color.Red("Error: " + err.Error())
		return "", "", "", "", "", err
	}
	shortcut, err := readString("Enter a shortcut for your command: ")
	if err != nil {
		color.Red("Error: " + err.Error())
		return "", "", "", "", "", err
	}

	return author, name, shortcut, command, help, nil
}
