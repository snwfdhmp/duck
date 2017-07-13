package main

import (
"fmt"
"os"
"bufio"
"./commands/init"
"./configuration"
)

func usageCommand(name string, desc string) {
	fmt.Printf("%s\t%s\n", name, desc)
}

func usage() {
	fmt.Println("usage : "+conf.APP_NAME+" <command>"+conf.END_STYLE+"\n")
	fmt.Println("Available commands :\n")
	fmt.Printf("command\tdescription\n")
	fmt.Printf("-------\t-----------\n")
	usageCommand("init", "init a new duck repo")
	usageCommand("config", "read a little bit of the config")
	usageCommand("console", "open duck console")

	//adds custom commands
	conf.Init()
	for _, val := range conf.Lang.Schemes {
		usageCommand(val.Label, val.Description)
	}
}

func RunCustomCmd(input string) {
	commands := conf.ParseCommand(input)

	//log number of commands
	//fmt.Println(len(commands), "commands")

	for _, cmd := range commands {
		output, _ := cmd.Output()
		fmt.Print(string(output))
	}

}

func Console() {
	var input string
	reader := bufio.NewReader(os.Stdin)
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

func CommandHandler(cmd string) {
	switch(cmd) {
	case "init":
		InitCmd.Run()
		break
	case "config":
		if(len(os.Args) < 3) {
			fmt.Println("Not enough arguments")
			os.Exit(1)
		}
		conf.Run(os.Args[2])
		break
	case "console":
		conf.Init()
		Console()
		break
	case "quit":
		fmt.Println(conf.YELLOW+"See you soon"+conf.END_STYLE)
	default:
		RunCustomCmd(cmd)
		// fmt.Printf(conf.RED+"unknown command: '%s'\n"+conf.END_STYLE, os.Args[1])
		// usage()
		// os.Exit(1);
		break
	}
}

func main() {
	if(len(os.Args) < 2) {
		usage()
		os.Exit(1)
	}
	CommandHandler(os.Args[1])
}