package cmd

import (
	"github.com/snwfdhmp/duck/pkg/data"
	"github.com/snwfdhmp/duck/pkg/pkg"
	"github.com/snwfdhmp/duck/pkg/projects"
	"github.com/spf13/cobra"
)

var (
	getOptions struct {
		Global  bool
		NoCheck bool
		Force   bool
	}
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a package from the internet",
	Long: `duck get <package> looks for <package> in your configured
repositories and download it from the first repository (based on
repositories array index) where its available.

Package are named following this pattern : 'author/name' (ie: 'snwfdhmp/go')
To try: duck get snwfdhmp/go`,
	Run: func(cmd *cobra.Command, args []string) {
		var target string

		if getOptions.Global {
			target = data.PackagesPath
		} else {
			target = projects.PackagesPath
		}

		for i := 0; i < len(args); i++ {
			pkg.Download(target, args[i])
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&getOptions.Force, "force", "f", false, "replace package if existing")
	getCmd.Flags().BoolVarP(&getOptions.Global, "global", "g", false, "install package for user instead of project")
	getCmd.Flags().BoolVar(&getOptions.NoCheck, "no-check", false, "skip file/folder existance checking")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
