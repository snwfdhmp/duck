package parser

import (
"../configuration"
"os/exec"
"strings"
"regexp"
"fmt"
"os"
)

type Tag struct {
	Label string
	Value string
}

var Tags []Tag

/**
 * @brief      Take the command string as input and return
 *             	the list of commands to be executed, after
 *             	having replaced the tags (like "$main", or
 *             	"$path") to their corresponding values
 *
 * @param      label  The command string (ie : "compile" or
 *                      "run")
 *
 * @return     { An array of shell command objects }
 */
func GetCommandArrFromInput(label string) []*exec.Cmd {
	//read config files
	conf.Init()

	//look for command in schemes
	content := conf.GetScheme(label)

	//create a slice that will contain the final
	// list of commands
	var commands []*exec.Cmd

	for _, cmd := range content {
		cmd = ParseTags(cmd)
		cmd = ParseDollarParams(cmd)

		arr := SplitCommand(cmd)

		//logging
		//fmt.Println(len(arr), arr)

		//create the command
		tmp := exec.Command(arr[0], arr[1:]...)

		//append it to the slice
		commands = append(commands, tmp)
	}
	return commands
}

func InitTags() {
	//declare each Tag
	Tags = []Tag {
		Tag {
			Label : "$main",
			Value : conf.GetMainPath(),
		},
		Tag {
			Label : "$path",
			Value : conf.GetProjectRoot(),
		},
		Tag {
			Label : "$name",
			Value : conf.GetName(),
		},
	}
}

//replace tags in command
func ParseTags(command string) string {
	InitTags()
	for _, tag := range Tags {
		command = strings.Replace(command, tag.Label, tag.Value, -1)
	}

	return command
}

//iterates through array to replace $1..$9 with real $1..$9
//	(like shell)
func ParseDollarParams(command string) string {
	for i := 1; i <= 9; i++ {
		sel := fmt.Sprintf("$%d", i)
		if(strings.Index(command, sel) != -1) {
			command = strings.Replace(command, sel, os.Args[i+1], -1)
		}
	}

	return command
}

func SplitCommand(command string) []string {
	//split into array using regexp (to let quoted string be 1 arg, as in shell)
	delimeter := "[^\\s\"']+|\"([^\"]*)\"|'([^']*)'"
	reg := regexp.MustCompile(delimeter)
	arr := reg.FindAllString(command, -1)

	return arr
}