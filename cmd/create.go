package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/snwfdhmp/duck/pkg/pkg"
	"github.com/snwfdhmp/duck/pkg/projects"
	"github.com/spf13/cobra"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func readArgsPositionally(args, names []string) map[string]string {
	res := make(map[string]string, 1)

	for i, n := range names {
		if len(args) <= i {
			res[n] = ""
			continue
		}
		res[n] = args[i]
	}
	return res
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a ling",
	Long: `Permits you to create a new ling (duck command).
You can use it in two modes :
    - interactive: 'duck create'
    - one line: 'duck create <author> <command_name> <shortcut> <commands_to_execute> <help_message>'`,
	Run: func(cmd *cobra.Command, args []string) {
		var command map[string]string
		command = make(map[string]string, 1)

		if len(args) == 0 { //if no args, run interactive mode
			fmt.Println("You are going to create a ling")

			command, err = ask(map[string]string{
				"Enter the author name":                                 "author",
				"Choose the command name (will be used as 'duck name')": "name",
				"Enter the command to execute":                          "cmd",
				"Enter message to be displayed for help":                "help",
				"Enter a shortcut for your command":                     "shortcut",
			})

			if err != nil {
				color.Red("Error: " + err.Error())
				return
			}
		} else if len(args) != 5 {
			color.Red("Syntax: duck create <author> <command_name> <shortcut> <commands_to_execute> <help_message>")
			return
		} else { //if 5 args, run one-line mode
			command = readArgsPositionally(args, []string{"author", "name", "shortcut", "cmd", "help"})
		}

		pkgName := filepath.Join(command["author"], "custom")
		err = pkg.Create(projects.PackagesPath, pkgName, command)
		if err != nil {
			color.Red("Error: " + err.Error())
			return
		}

		color.Green("Ling successfully created")
	},
}

func readString(text string) (string, error) {
	fmt.Print(text + ": ")
	input, err := reader.ReadString('\n')
	return input[:len(input)-1], err
}

func init() {
	RootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ask(list map[string]string) (map[string]string, error) {
	answers := make(map[string]string, 1)
	var err error

	for text, dst := range list {
		answers[dst], err = readString(text)
		if err != nil {
			return nil, err
		}
	}
	return answers, nil
}
