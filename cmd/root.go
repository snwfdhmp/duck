package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var cfgFile string

var (
	err error
	fs  afero.Fs = afero.NewOsFs()

	Yellow    = color.New(color.FgYellow).Sprint
	Blue      = color.New(color.FgCyan).Sprint
	Green     = color.New(color.FgGreen).Sprint
	Red       = color.New(color.FgRed).Sprint
	Magenta   = color.New(color.FgMagenta).Sprint
	repoColor = color.New(color.FgCyan).Sprint
	pkgColor  = color.New(color.FgYellow).Sprint
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "duck",
	Short: "Developer assistant",
	Long: `Duck is a developer assistant, it helps you to automate task
by creating custom commands (called 'lings'), or by managing
your projects.`,
	Example: `'go install && go run test.go ' => 'duck build'
'git add * && git commit -m "My commit text" && git push' => 'duck commit "My commit text"'`,
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
}
