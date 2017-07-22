/*
	@TODO describe parser with tags
*/
package parser

import (
	"fmt"
	"github.com/snwfdhmp/duck/pkg/configuration"
	"github.com/snwfdhmp/duck/pkg/logger"
	"os/exec"
	"regexp"
	"strings"
)

type Tag struct {
	Label string
	Value string
}

var Tags []Tag

//GetCommandsArrFromInput Take the command string as input and return
//the list of commands to be executed, after
//having replaced the tags (like "$main", or
//"$path") to their corresponding values
//
//@param      label  The command string (ie : "compile" or
//                   "run")
//
//@return     { An array of shell command objects }
func GetCommandArrFromInput(label ...string) []*exec.Cmd {
	//read config files
	conf.Init()

	//look for command in schemes
	content := conf.GetCommands(label[0])

	//create a slice that will contain the final
	// list of commands
	var commands []*exec.Cmd

	for _, cmd := range content {
		shouldContinue := true
		cmd = ParseTags(cmd)
		cmd, shouldContinue = ParseDollarParams(cmd, label[1:]...)
		if shouldContinue != true {
			return []*exec.Cmd{}
		}
		//logging
		//fmt.Println(len(arr), arr)

		arrCmd := SplitCommand(cmd)
		if arrCmd[0] == "@" {
			toAdd := GetCommandArrFromInput(arrCmd[1:]...)
			for _, tmp := range toAdd {
				commands = append(commands, tmp)
			}
			continue
		}

		//create the command
		tmp := exec.Command("sh", "-c", cmd)

		//append it to the slice
		commands = append(commands, tmp)
	}
	return commands
}

//InitTags declares the tags being parsed
func InitTags() {
	//declare each Tag
	Tags = []Tag{
		Tag{
			Label: "$main",
			Value: conf.GetMainPath(),
		},
		Tag{
			Label: "$path",
			Value: conf.GetProjectRoot(),
		},
		Tag{
			Label: "$name",
			Value: conf.GetName(),
		},
	}
}

//ParseTags replace tags in command
func ParseTags(command string) string {
	InitTags()
	for _, tag := range Tags {
		command = strings.Replace(command, tag.Label, tag.Value, -1)
	}

	return command
}

//ParseDollarParams iterates through array to replace $1..$9 with real $1..$9
//	(like shell)
func ParseDollarParams(command string, args ...string) (string, bool) {
	for i := 1; i <= 9; i++ {
		sel := fmt.Sprintf("$%d", i)
		if strings.Index(command, sel) != -1 {
			if len(args) < i {
				fmt.Printf(logger.BLUE+"$%d "+logger.RED+"argument was not provided.\n"+logger.END_STYLE, i)
				return "", false
			}
			command = strings.Replace(command, sel, args[i-1], -1)
		}
	}

	return command, true
}

//SplitCommand splits the command into an array
//corresponding to the shell args
func SplitCommand(command string) []string {
	//split into array using regexp (to let quoted string be 1 arg, as in shell)
	delimeter := "[^\\s\"']+|\"([^\"]*)\"|'([^']*)'"
	reg := regexp.MustCompile(delimeter)
	arr := reg.FindAllString(command, -1)

	for i, arg := range arr {

		//delete extremities quotes
		// ex : git, commit, -m, "Message to be displayed"
		//   => git, commit, -m, Message to be displayed
		//
		// it avoids extra quotes when the argument is exported
		//  into other files or services (such as git)
		if arg[0] == '"' && arg[len(arg)-1] == '"' {
			arg = arg[1 : len(arg)-1]
			arr[i] = arg
		}
		//arr[i] = strings.Replace(arg, "\"", "", -1)
	}
	return arr
}
