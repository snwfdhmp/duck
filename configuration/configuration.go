package conf

import (
"fmt"
"encoding/json"
"os"
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

//@todo rename Schemes to Ducklings ("caneton" in French (duck's children))
type Scheme struct {
	Label string
	Commands []string
	Description string
	Aliases []string
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
	//looking for the commands corresponding to the label
	for _, val := range Lang.Schemes {
		if(val.Label == label) { // if scheme's label is input, return it
			return val.Commands
		} else { //else look in its aliases
			for _, alias := range val.Aliases {
				if(alias == label) {
					return val.Commands
				}
			}
		}
	}

	//if nothing found, return an error
	return []string {"echo This command doesn't exists"} //@todo handle errors better
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


func checkErr(err error) {
	if (err != nil) {
		panic(err)
	}
}