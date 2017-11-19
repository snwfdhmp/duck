package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/snwfdhmp/duck/pkg/data"
	"github.com/snwfdhmp/duck/pkg/projects"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List projects, has sub-commands",
	Run: func(cmd *cobra.Command, args []string) {
		for _, p := range data.Projects.Keys() {
			fmt.Println("-", Blue(p.Name()), "=>", Yellow(p.Value()))
		}

		fmt.Println(len(data.Projects.Keys()), "projects.")
	},
}

func RepairProject(path string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return
	}
	projects.InitProject([]string{}, absolutePath, false, true)
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
