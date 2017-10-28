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

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// repairCmd represents the repair command
var repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			wd, err := os.Getwd()
			if err != nil {
				color.Red("Cannot get working directory :" + err.Error())
				return
			}
			args = append(args, wd)
		}

		if len(args) == 1 {
			RunRepair(args[0])
		} else {
			for _, path := range args {
				fmt.Println("-", path)
				RunRepair(path)
			}
		}
	},
}

func RunRepair(path string) {
	healthy, err := IsHealthy(path)
	if err != nil {
		color.Red("Doctor cannot work :" + err.Error())
		return
	}

	if healthy {
		color.Green("Project is already healthy.")
		return
	}

	RepairProject(path)
	healthy, err = IsHealthy(path)
	if err != nil {
		color.Red("Doctor verify repair worked :" + err.Error())
		return
	}
	if healthy {
		color.Green("This project has been repaired !")
	} else {
		color.Red("I failed to repaired this project.")
	}
}

func init() {
	RootCmd.AddCommand(repairCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repairCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repairCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
