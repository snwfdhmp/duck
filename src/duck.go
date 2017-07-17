package main

import (
	// "./commands/init"
	"./configuration"
	"./parser"
	"./usage"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//@todo add support of (args ...string)
//  for command handlers to enhance
//  the duck console usage

const (
	DUCK_VERSION = "a1.0.0"
)

/**
 * Execute a user custom's command
 * @param {string} input [the user input]
 */
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
		checkErr(err)
		stderr, err := cmd.StderrPipe()
		checkErr(err)

		err = cmd.Start()
		checkErr(err)

		//print stdout and stderr
		output, err := ioutil.ReadAll(stdout)
		checkErr(err)
		slurp, err := ioutil.ReadAll(stderr)
		checkErr(err)
		fmt.Print(conf.RED + string(slurp) + conf.END_STYLE)
		fmt.Print(conf.GREEN + string(output) + conf.END_STYLE)

		cmd.Wait()
	}

}

/**
 * The console will loop on stdin until
 *  the user inputs "quit"
 */
func Console() {
	var input string                    //will contain input from stdin
	reader := bufio.NewReader(os.Stdin) //reader initialized for stdin

	for input != "quit" {
		//read input
		fmt.Print(conf.APP_NAME + ":" + conf.GetName() + "> ")
		input, _ = reader.ReadString('\n')

		//delete the '\n'
		input = input[:len(input)-1]

		//throw error for special cases
		if input == "config" {
			fmt.Println("Not available in console mode yet.")
			continue
		}

		command := parser.SplitCommand(input)

		//handle input, break if needed
		if !CommandHandler(command...) {
			break
		}
	}
}

/**
 * Will route any command supported by duck or custom conf
 *  to the function that handles it
 * @param {string} cmd 			[the command asked]
 * @todo better args
 */
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
		break
	case "config": //print a config property (@todo add command to modify
		if len(cmd) < 2 {
			fmt.Println("Not enough arguments")
			os.Exit(1)
		}
		conf.Run(cmd[1])
		break
	case "console": //launch duck console
		conf.Init()
		Console()
		break
	case "lings":
		conf.PrintPackages()
		break
	case "help": //prints help
		usage.PrintAll()
		break
	case "repo-list":
		conf.PrintRepos()
		break
	case "install":
		if len(cmd) >= 2 {
			conf.InstallPkg(cmd[1:]...)
		} else {
			conf.InstallPkg()
		}
		break
	case "uninstall":
		if len(cmd) >= 2 {
			conf.UninstallDuckling(cmd[1:]...)
		} else {
			fmt.Println(conf.RED + "Not enough arguments" + conf.END_STYLE)
		}
		break
	case "man": //prints manual
		usage.Man()
		break
	case "version": //print duck version
		fmt.Println(conf.APP_NAME, DUCK_VERSION)
		break
	case "quit": //quit
		fmt.Println(conf.BLUE + "See you soon" + conf.END_STYLE)
		shouldBreak = true
		break
	default: //if input is none of the "general" commands, use custom ones
		RunCustomCmd(cmd...)
		break
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
