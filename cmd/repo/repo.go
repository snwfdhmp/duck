package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/snwfdhmp/duck/pkg/data"
	"github.com/spf13/cobra"
)

const (
	DefaultRepoName = "core"
	DefaultRepoURL  = "http://raw.githubusercontent.com/snwfdhmp/duck-core/master/"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Get list of your repositories",
	Run: func(cmd *cobra.Command, args []string) {
		repoColor := color.New(color.FgCyan).SprintFunc()
		urlColor := color.New(color.FgYellow).SprintFunc()

		repos, err := getRepos()
		if err != nil {
			color.Red("Could not load repos : " + err.Error())
			return
		}
		for name, url := range repos {
			fmt.Println("-", repoColor(name), "=>", urlColor(url))
		}
	},
}

func init() {
	RootCmd.AddCommand(repoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
