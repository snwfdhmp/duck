package conf

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	ProjectRoot  string
	Ducklings    []string
	Name         string
	Lang         string
	VersionMajor string
	VersionMinor string
	Env          string
	Main         string
}

//@todo rename Schemes to Ducklings ("caneton" in French (duck's children))
type Duckling struct {
	Label       string
	Commands    []string
	Description string
	Aliases     []string
}

type Duckfile struct {
	Ducklings []Duckling
}

var (
	Conf        Configuration
	Ducklings   []Duckling
	verboseMode bool
)

const (
	RED    = "\033[1;31m"
	GREEN  = "\033[1;32m"
	YELLOW = "\033[1;33m"
	BLUE   = "\033[1;36m"

	ITALIC = "\033[3m"

	END_STYLE = "\033[0m"

	APP_NAME = YELLOW + "duck" + END_STYLE
)

//@todo: delete that shit and but add a similar func to parser
func Run(command string) {
	Init()
	switch command {
	case "name":
		fmt.Println(GetName())
		break
	default:
		fmt.Printf("No handler for command '%s'\n", command)
		break
	}
}

/**
 * @brief      Reads the configuration files
 *
 * @return     bool: status (ok|not ok)
 */
func Init() bool {
	dir, _ := os.Getwd()

	if !ExistsConfIn(dir) {
		fmt.Println("No duck repo found in", dir)
		return false
	}

	LoadProjectConfig(dir)

	LoadDuckfiles()

	return true
}

//load a JSON file into its correctly typed interface
func LoadFileJson(path string, object interface{}) bool {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		fmt.Println("Error: not found", path, "found")
		return false
	}
	checkErr(err)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(object)
	checkErr(err)
	return true
}

//load project conf file into Conf
func LoadProjectConfig(dir string) bool {
	path := dir + "/.duck/project.conf"
	LoadFileJson(path, &Conf)
	//fmt.Println("Loaded conf for project", Conf.Name, "in language", Conf.Lang)
	return true
}

func BuildDuckfilePath(str string) string {
	return Conf.ProjectRoot + "/.ducklings/" + str + ".duckling"
}

func verbose(str string) {
	if verboseMode {
		fmt.Println(str)
	}
}

//load LangFile (@todo "duckling") to Lang
func LoadDuckfiles() {
	Ducklings = []Duckling{}
	for _, duckling := range Conf.Ducklings {
		path := BuildDuckfilePath(duckling)
		verbose(YELLOW + "from " + duckling + END_STYLE)
		var Duckfile Duckfile
		LoadFileJson(path, &Duckfile)
		for _, tmp := range Duckfile.Ducklings {
			verbose(BLUE + " importing " + duckling + "/" + tmp.Label + END_STYLE)
			Ducklings = append(Ducklings, tmp)
		}
	}
}

func PrintDucklings() {
	Init()
	for _, duckling := range Conf.Ducklings {
		fmt.Println(duckling)
		count++
	}
	fmt.Println("total:", count)
	verboseMode = true
	LoadDuckfiles()
	verboseMode = false
}

func GetCommands(label string) []string {
	//looking for the commands corresponding to the label
	for _, val := range Ducklings {
		if val.Label == label { // if scheme's label is input, return it
			return val.Commands
		} else { //else look in its aliases
			for _, alias := range val.Aliases {
				if alias == label {
					return val.Commands
				}
			}
		}
	}

	//if nothing found, return an error
	return []string{"echo " + RED + "Unknown command '" + label + "'" + END_STYLE} //@todo handle errors better
}

func ExistsConfIn(dir string) bool {
	DUCK_DIR := ".duck"

	duckPath := dir + "/" + DUCK_DIR
	if _, err := os.Stat(duckPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
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
	if err != nil {
		panic(err)
	}
}

func AskConf() {
	var NewConf Configuration

	NewConf.Name = askProperty("Project name")
	fmt.Println("Name :", NewConf.Name)
}

func getRidNewLine(str string) string {
	return str[:len(str)-1]
}

func askProperty(prompt string) string {
	reader := bufio.NewReader(os.Stdin) //reader initialized for stdin
	fmt.Print(prompt)
	tmp, _ := reader.ReadString('\n')
	return getRidNewLine(tmp)
}
