package main

import (
"fmt"
"os"
"bufio"
"./commands/init"
"./configuration"
)


const (
	DUCK_VERSION = "b0.1"
)

/**
 * This function defines the command's usage template
 *  (every command usage should be printed using this
 *  function in order to be correctly displayed)
 * @param  {string} name    	[name of the command]
 * @param  {string} desc    	[description of the command]
 */
func usageCommand(name string, desc string) {
	fmt.Printf("%s\t%s\n", name, desc)
}

/**
 * Prints duck usage
 */
func usage() {
	//print head
	fmt.Println("usage : "+conf.APP_NAME+" <command>"+conf.END_STYLE+"\n")
	fmt.Println("Available commands :\n")
	fmt.Println("command\tdescription")
	fmt.Println("-------\t-----------")

	//print general commands
	usageCommand("init", "init a new duck repo")
	usageCommand("config", "read a little bit of the config")
	usageCommand("console", "open duck console")
	usageCommand("version", "print duck's version")

	//print custom commands
	conf.Init()
	for _, val := range conf.Lang.Schemes {
		usageCommand(val.Label, val.Description)
	}
}

/**
 * Execute a user custom's command
 * @param {string} input [the user input]
 */
func RunCustomCmd(input string) {
	//get commands array from <lang>.duck
	commands := conf.ParseCommand(input)

	//log number of commands
	//fmt.Println(len(commands), "commands")

	for _, cmd := range commands {
		output, _ := cmd.Output()
		fmt.Print(string(output))
	}

}

/**
 * The console will loop on stdin until
 *  the user inputs "quit"
 */
func Console() {
	var input string //will contain input from stdin
	reader := bufio.NewReader(os.Stdin) //reader initialized for stdin

	for (input != "quit") {
		//read input
		fmt.Print(conf.APP_NAME+":"+conf.GetName()+"> ")
		input, _ = reader.ReadString('\n')

		//delete the '\n'
		input = input[:len(input)-1]

		//throw error for special cases
		if(input == "config") {
			fmt.Println("Not available in console mode yet.")
			continue
		}

		//handle input
		CommandHandler(input)
	}
}

/**
 * Will route any command supported by duck or custom conf
 *  to the function that handles it
 * @param {string} cmd 			[the command asked]
 */
func CommandHandler(cmd string) {
	//managing shortcuts
	if(cmd == "sh" || cmd == "shell") {
		cmd = "console"
	} else if (cmd == "c"){
		cmd = "compile"
	}

	//handling command
	switch(cmd) {
	case "init": //init a new duck repo
		InitCmd.Run()
		break
	case "config": //print a config property @todo add command to modify
		if(len(os.Args) < 3) {
			fmt.Println("Not enough arguments")
			os.Exit(1)
		}
		conf.Run(os.Args[2])
		break
	case "console": //launch duck console
		conf.Init()
		Console()
		break
	case "version": //print duck version
		fmt.Println("duck", DUCK_VERSION)
		break
	case "quit": //quit
		fmt.Println(conf.BLUE+"See you soon"+conf.END_STYLE)
		break
	default: //if input is none of the "general" commands, use custom ones
		RunCustomCmd(cmd)
		break
	}
}

/**
 * main function, execution flow will start here
 */
func main() {
	//if no args, print usage and exit with error
	if(len(os.Args) < 2) {
		usage()
		os.Exit(1)
	}
	//give control to CommandHandler
	CommandHandler(os.Args[1])
}