package usage

import (
"fmt"
"../configuration"
)

/**
 * CommandUsage type
 */
type CmdUsg struct {
	Name string
	Desc string
	Aliases []string
}

var (
	Commands []CmdUsg
)

func Load() {
	//duck's native commands
	Commands = []CmdUsg {
		CmdUsg {
			Name:"init",
			Desc:"init a new duck repo",
			},
		CmdUsg {
			Name:"config",
			Desc:"read a little bit of the config",
			},
		CmdUsg {
			Name:"console",
			Desc:"open duck console",
			},
		CmdUsg {
			Name:"version",
			Desc:"print duck's version",
			},
		CmdUsg {
			Name:"help",
			Desc:"print duck's help",
			},
		CmdUsg {
			Name:"man",
			Desc:"print duck's extended help",
			},
	}

	//custom commands
	conf.Init()
	var tmp CmdUsg
	for _, val := range conf.Lang.Schemes {
		tmp = CmdUsg {
			Name:val.Label,
			Desc:val.Description,
			Aliases:val.Aliases,
		}
		Commands = append(Commands, tmp)
	}
}

/**
 * Prints duck usage
 */
func PrintAll() {
	Load()
	//print head
	fmt.Println("usage : "+conf.APP_NAME+" <command>"+conf.END_STYLE+"\n")
	fmt.Println("Available commands :\n")
	fmt.Println("command\t\tdescription")
	fmt.Println("-------\t\t-----------")

	for _, cmd := range Commands {
		printCommand(cmd)
	}
}

/**
 * This function defines the command's usage template
 *  (every command usage should be printed using this
 *  function in order to be correctly displayed)
 * @param  {string} name    	[name of the command]
 * @param  {string} desc    	[description of the command]
 */
func printCommand(cmd CmdUsg) {
	fmt.Printf("%s\t", cmd.Name)
	fmt.Printf("\t%s\n", cmd.Desc)
}

func Man() {
	Load()
	fmt.Println(conf.APP_NAME, "help:")
	for _, cmd := range Commands {
		ManCommand(cmd)
	}
}

func ManCommand(cmd CmdUsg) {
	fmt.Print("\n")
	fmt.Println("-",cmd.Name)

	fmt.Print("...Aliases:")
	for _, alias := range cmd.Aliases {
		fmt.Printf(" %s", alias)
	}
	fmt.Print("\n")

	fmt.Println("...Description:", cmd.Desc)
}