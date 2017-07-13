package conf

import (
"fmt"
"encoding/json"
"os"
"os/exec"
"strings"
"regexp"
)

type Configuration struct {
	ProjectRoot string
	Name string
	Lang string
	VersionMajor string
	VersionMinor string
	Env string
	Main string
}

type Scheme struct {
	Label string
	Commands []string
	Description string
}

type LangFile struct {
	Schemes []Scheme
}

var (
	Conf Configuration
	Lang LangFile
)

const (
	RED="\033[1;31m"
	GREEN="\033[1;32m"
	YELLOW="\033[1;33m"
	BLUE="\033[1;36m"

	ITALIC="\033[3m"

	END_STYLE="\033[0m"

	APP_NAME=YELLOW+"duck"+END_STYLE
)

func Run(command string) {
	Init()
	switch(command) {
	case "lang":
		fmt.Println(GetLang())
		break;
	case "name":
		fmt.Println(GetName())
		break;
	default:
		fmt.Printf("No handler for command '%s'\n", command)
		break;
	}
}

func Init() bool {
	dir, _ := os.Getwd()

	if(!ExistsConfIn(dir)) {
		fmt.Println("No duck repo found in", dir)
		return false
	}

	LoadProjectConfig(dir)

	LoadLangFile(dir, GetLang())

	return true
}

func LoadFileJson(path string, object interface{}) bool{
	file, err := os.Open(path)
	if(os.IsNotExist(err)) {
		fmt.Println("Error: not found", path, "found")
		return false
	}
	checkErr(err)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(object)
	checkErr(err)
	return true
}

func LoadProjectConfig(dir string) bool{
	path := dir+"/.duck/project.conf"
	LoadFileJson(path, &Conf)
	//fmt.Println("Loaded conf for project", Conf.Name, "in language", Conf.Lang)
	return true
}

func LoadLangFile(dir string, lang string) {
	path := dir+"/.duck/"+lang+".duck"
	LoadFileJson(path, &Lang)
}

func GetScheme(label string) []string{
	for _, val := range Lang.Schemes {
		if(val.Label == label) {
			return val.Commands
		}
	}
	return []string {"echo"} //@todo handle errors
}

func GetLang() string {
	return Conf.Lang
}

func GetName() string {
	return Conf.Name
}

func GetMain() string {
	return Conf.Main
}

func GetProjectRoot() string {
	return Conf.ProjectRoot
}

func GetMainPath() string {
	return Conf.ProjectRoot + "/" + Conf.Main
}

func ParseCommand(label string) []*exec.Cmd {
	//read config files
	Init()

	//look for command in schemes
	content := GetScheme(label)

	var commands []*exec.Cmd

	for _, cmd := range content {
	//replace tags in command
		cmd = strings.Replace(cmd, "$main", GetMainPath(), -1)
		cmd = strings.Replace(cmd, "$path", GetProjectRoot(), -1)
		cmd = strings.Replace(cmd, "$name", GetName(), -1)

		for i := 1; i <= 9; i++ {
			sel := fmt.Sprintf("$%d", i)
			if(strings.Index(cmd, sel) != -1) {
				cmd = strings.Replace(cmd, sel, os.Args[i+1], -1)
			}
		}

		//split into array using regexp (to let quoted string be 1 arg, as in shell)
		delimeter := "[^\\s\"']+|\"([^\"]*)\"|'([^']*)'"
    	reg := regexp.MustCompile(delimeter)
    	arr := reg.FindAllString(cmd, -1)

		//logging
		//fmt.Println(len(arr), arr)

		tmp := exec.Command(arr[0], arr[1:]...)
		commands = append(commands, tmp)
	}
	return commands
}

func ExistsConfIn(dir string) bool{
	DUCK_DIR := ".duck"

	duckPath := dir+"/"+DUCK_DIR
	if _, err := os.Stat(duckPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}


func checkErr(err error) {
	if (err != nil) {
		panic(err)
	}
}