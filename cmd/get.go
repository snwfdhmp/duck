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
	"io"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//add -f to force creation of dir
var nocheck bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()

		err := loadProjectConfig()
		if err != nil {
			color.Red("Unable to load project configuration.")
			color.Red("Error: " + err.Error())
			return
		}

		pkgs, err := projectCfg.GetSection("packages")
		if err != nil {
			color.Red("Unable to get 'packages' section into configuration file")
			color.Red("Error: " + err.Error())
			return
		}

		path, err := pkgs.GetKey("directory")
		if err != nil {
			color.Red("Unable to get the 'directory' key from 'packages' section into configuration file")
			color.Red("Error: " + err.Error())
			return
		}

		pkgsPath := ".duck/" + path.String()

		for i := 0; i < len(args); i++ {
			// if args[i][len(args[i])-1] == "/" { //delete '/' if in last position, should be tested before use
			// 	args[i] = args[i][:len(args[i])-1]
			// }
			arr := strings.Split(args[i], "/")
			var out afero.File
			currentPath := pkgsPath
			for j := 0; j < len(arr); j++ {
				currentPath += "/" + arr[j]
				if j < len(arr)-1 {
					if !nocheck {
						exists, err := afero.Exists(fs, currentPath)
						if err != nil {
							color.Red("Could not test whether '" + currentPath + "' exists or not. (not implemented: use --no-check to force)")
							color.Red("Error: " + err.Error())
							return
						}
						if exists {
							continue
						}
					}
					err = fs.Mkdir(currentPath, 0777)
					if err != nil {
						color.Red("Could not create '" + currentPath)
						color.Red("Error: " + err.Error())
						if !nocheck {
							return
						}
					}
				} else {
					if !nocheck {
						currentPath += ".duckpkg.ini"
						exists, err := afero.Exists(fs, currentPath)
						if err != nil {
							color.Red("Could not test whether '" + currentPath + "' exists or not. (not implemented: use --no-check to force)")
							color.Red("Error: " + err.Error())
							return
						}
						if exists && !force {
							color.Red("This package seems to be already installed. Use -f to install over")
							return
						}
					}
					out, err = fs.Create(currentPath)
					if err != nil {
						color.Red("Could not create '" + currentPath)
						color.Red("Error: " + err.Error())
						if !nocheck {
							return
						}
					}
				}
			}
			defer out.Close()
			repos := viper.GetStringMap("repos")
			repoColor := color.New(color.FgCyan).Sprint
			pkgColor := color.New(color.FgYellow).Sprint
			for name, url := range repos {
				pkgUrl := fmt.Sprintf("%s%s.duckpkg.ini", url, args[i])
				resp, err := http.Get(pkgUrl) //test errors
				if err != nil {
					color.Red("Could not download from '" + pkgUrl + "'")
					color.Red("Error: " + err.Error())
					return
				}
				defer resp.Body.Close()

				_, err = io.Copy(out, resp.Body)

				color.Green("Successfully installed '" + pkgColor(args[i]) + color.New(color.FgGreen).Sprint("' from ") + repoColor(name))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&force, "force", "f", false, "replace package if existing")
	getCmd.Flags().BoolVar(&nocheck, "no-check", false, "skip file/folder existance checking")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
