package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/snwfdhmp/duck/pkg/projects"
	"github.com/spf13/cobra"
)

// repairCmd represents the repair command
var repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "Repair project",
	Long: `Repair project.

Usage: duck repair <dir1> <dir2> ...
Default : current working directory`,
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
	healthy, err := projects.IsHealthy(path)
	if err != nil {
		color.Red("Doctor cannot work :" + err.Error())
		return
	}

	if healthy {
		color.Green("Project is already healthy.")
		return
	}

	RepairProject(path)
	healthy, err = projects.IsHealthy(path)
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
	//RootCmd.AddCommand(repairCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repairCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repairCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
