package usage

import (
"fmt"
"../configuration"
)

type CommandDescriptor struct {
	Name string
	Desc string
}

/**
 * Prints duck usage
 */
func PrintAll() {
	//print head
	fmt.Println("usage : "+conf.APP_NAME+" <command>"+conf.END_STYLE+"\n")
	fmt.Println("Available commands :\n")
	usageCommand("command","description")
	usageCommand("-------","-----------")

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
 * This function defines the command's usage template
 *  (every command usage should be printed using this
 *  function in order to be correctly displayed)
 * @param  {string} name    	[name of the command]
 * @param  {string} desc    	[description of the command]
 */
func usageCommand(name string, desc string) {
	fmt.Printf("%s\t\t%s\n", name, desc)
}
