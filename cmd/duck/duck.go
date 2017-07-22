package main

import (
	"bufio"
	"fmt"
	"github.com/snwfdhmp/duck/pkg/configuration"
	"github.com/snwfdhmp/duck/pkg/logger"
	"github.com/snwfdhmp/duck/pkg/parser"
	"github.com/snwfdhmp/duck/pkg/usage"
	"io/ioutil"
	"os"
)

//@todo add support of (args ...string)
//  for command handlers to enhance
//  the duck console usage
const (
	//DuckVersion current tool version
	DuckVersion = "v0.1.2"
)

var (
	log logger.Logger
)

//RunCustomCmd executes a user custom's command
//
//@param {string} input [the user input]
func RunCustomCmd(input ...string) {
	//get commands array from <lang>.duck
	commands := parser.GetCommandArrFromInput(input...)

	//log number of commands
	//fmt.Println(len(commands), "commands")

	for _, cmd := range commands {
		/**
		 * pipe stdout and stderr
		 * to handle error nicely
		 * and being able to print
		 * command errors to user
		 */
		stdout, err := cmd.StdoutPipe()
		log.Err(err, "Failed to pipe stdout")

		stderr, err := cmd.StderrPipe()
		log.Err(err, "Failed to pipe stderr")

		err = cmd.Start()
		log.Fatal(err, "Failed to start command")

		//read stdout and stderr
		output, err := ioutil.ReadAll(stdout)
		log.Err(err, "Failed to read stdout")

		slurp, err := ioutil.ReadAll(stderr)
		log.Err(err, "Failed to read stderr")

		//print stdout and stderr
		fmt.Print(logger.RED + string(slurp) + logger.END_STYLE)
		fmt.Print(logger.GREEN + string(output) + logger.END_STYLE)

		cmd.Wait()
	}

}

//Console will loop on stdin until
//the user inputs "quit"
func Console() {
	var input string                      //will contain input from stdin
	scanner := bufio.NewScanner(os.Stdin) //scanner initialized for stdin

	for input != "quit" {
		//read input
		fmt.Print(logger.APP_NAME + ":" + conf.GetName() + "> ")
		scanner.Scan()
		input := scanner.Text()

		//throw error for special cases
		if input == "config" {
			fmt.Println("Not available in console mode yet.")
			continue
		}

		command := parser.SplitCommand(input) //parse command

		//handle input, break if needed
		if !CommandHandler(command...) {
			break
		}
	}
}

//CommandHandler will route any command supported by duck or custom conf
//to the function that handles it
//It takes as param the cmd to run and its params as ellipsis
//
//Example : "CommandHandler("install", "-f", "snwfdhmp/std", "snwfdhmp/go")
func CommandHandler(cmd ...string) bool {
	shouldBreak := false //should we stop execution ?

	//managing shortcuts
	if cmd[0] == "sh" || cmd[0] == "shell" {
		cmd[0] = "console"
	} else if cmd[0] == "q" {
		cmd[0] = "quit"
	}

	//handling command
	switch cmd[0] {
	case "init": //init a new duck repo
		conf.AskConf()
	case "config": //print a config property (@todo add command to modify
		if len(cmd) < 2 {
			fmt.Println("Not enough arguments")
		} else {
			conf.Run(cmd[1])
		}
	case "console": //launch duck console
		conf.Init()
		Console()
	case "lings":
		conf.PrintPackages()
	case "help": //prints help
		usage.PrintAll()
	case "repo-list":
		conf.PrintRepos()
	case "install":
		if len(cmd) >= 2 {
			conf.InstallPkgs(cmd[1:]...)
		} else {
			conf.InstallPkgs()
		}
	case "uninstall":
		if len(cmd) >= 2 {
			conf.UninstallPkgs(cmd[1:]...)
		} else {
			fmt.Println(logger.RED + "Not enough arguments" + logger.END_STYLE)
		}
	case "repo-add":
		if len(cmd) != 3 {
			fmt.Println(logger.RED + "usage: @ repo-add <name> <url>" + logger.END_STYLE)
		} else {
			conf.AddRepo(cmd[1], cmd[2])
		}
	case "man": //prints manual
		usage.Man()
	case "version": //print duck version
		fmt.Println(logger.APP_NAME, DuckVersion) //actual tool version
	case "quit": //quit
		fmt.Println(logger.BLUE + "See you soon" + logger.END_STYLE)
		shouldBreak = true
	default: //if input is none of the "general" commands, use custom ones
		RunCustomCmd(cmd...)
	}

	return !shouldBreak
}

/**
 * main function, execution flow will start here
 */
func main() {
	//if no args, print usage and exit with error
	if len(os.Args) < 2 {
		usage.PrintAll()
		os.Exit(1)
	}
	//give control to CommandHandler
	CommandHandler(os.Args[1:]...)
}
