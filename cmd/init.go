package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/snwfdhmp/duck/pkg/projects"
	"github.com/spf13/cobra"
)

var (
	initOptions struct {
		Force bool
	}
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inits a new duck repo in the current directory",
	Long: `Inits a new duck repo in the current directory with structure :
 .duck/
   packages/
   conf.ini`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			color.Red(err.Error())
			return
		}

		projects.InitProject(args, wd, true, initOptions.Force)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&initOptions.Force, "force", "f", false, "Delete .duck if it exists")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
