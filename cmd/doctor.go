package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/snwfdhmp/duck/pkg/projects"
	"github.com/spf13/cobra"
)

var (
	doctorOptions struct {
		Repair bool
	}
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Gives project's status",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			wd, err := os.Getwd()
			if err != nil {
				color.Red("Cannot get working directory :" + err.Error())
				return
			}
			args = append(args, wd)
		}

		var toRun Funcs

		toRun.Add(RunDoctor)

		if doctorOptions.Repair {
			toRun.Add(RunRepair)
		}

		if len(args) == 1 {
			toRun.Run(args[0])
		} else {
			for _, path := range args {
				fmt.Println("-", path)
				toRun.Run(path)
			}
		}
	},
}

type Funcs []func(string)

func (f *Funcs) Add(fn func(string)) {
	*f = append(*f, fn)
}

func (f *Funcs) Run(s string) {
	for _, fn := range *f {
		fn(s)
	}
}

func RunDoctor(path string) {
	healthy, err := projects.IsHealthy(path)
	if err != nil {
		color.Red("Doctor cannot work :" + err.Error())
		return
	}

	if healthy {
		color.Green("You look good ! :-)")
	} else {
		color.Magenta("You should repair.")
	}
}

func init() {
	RootCmd.AddCommand(doctorCmd)
	doctorCmd.Flags().BoolVarP(&doctorOptions.Repair, "repair", "r", false, "will repair project if broken")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doctorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doctorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
