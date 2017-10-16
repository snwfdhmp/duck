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
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-ini/ini"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	fs         afero.Fs = afero.NewOsFs()
	projectCfg *ini.File
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "duck",
	Short: "Developer assistant",
	Long: `Duck is a developer assistant, it helps you to automate task
by creating custom commands (called 'lings'), or by managing
your projects.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.duck.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".duck" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".duck")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func loadProjectConfig() error {
	exists, err := afero.Exists(fs, ".duck/conf.ini")
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("'.duck/conf.ini' seems not to exist.")
	}

	projectCfg, err = ini.Load(".duck/conf.ini")
	if err != nil {
		return err
	}

	return nil
}

func getConfigString(str string) (string, error) {
	arr := strings.Split(str, ".")
	section, err := projectCfg.GetSection(arr[0])
	if err != nil {
		return "", err
	}
	key, err := section.GetKey(arr[1])
	if err != nil {
		return "", err
	}
	return key.String(), nil
}

func DuckGlobalConfPath() (string, error) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := home + "/.duck"

	exists, err := afero.Exists(fs, path)
	if err != nil {
		return "", err
	}
	if exists {
		return path, nil
	}

	err = fs.Mkdir(path, 0755)
	if err != nil {
		return "", err
	}

	return path, nil
}

func getDuckDataFromPath(path string) (*ini.File, error) {
	exists, err := afero.Exists(fs, path)
	if err != nil {
		return &ini.File{}, err
	}
	if !exists {
		_, err := fs.Create(path)
		if err != nil {
			return &ini.File{}, err
		}
	}

	config, err := ini.Load(path)
	if err != nil {
		return &ini.File{}, err
	}

	_, err = config.GetSection("projects")
	if err != nil {
		_, err = config.NewSection("projects")
		if err != nil {
			return &ini.File{}, err
		}
		err := config.SaveTo(path)
		if err != nil {
			return &ini.File{}, err
		}
	}

	return config, nil
}

func getDuckData() (*ini.File, error) {
	configPath, err := getDuckDataPath()
	if err != nil {
		return &ini.File{}, err
	}

	return getDuckDataFromPath(configPath)
}

func getDuckDataPath() (string, error) {
	path, err := DuckGlobalConfPath()

	return path + "/data.ini", err
}
