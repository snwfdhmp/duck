package conf

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

type Repo struct {
	Name string
	URL  string
}

type AppConfiguration struct {
	Repos []Repo
}

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
	Dependencies []string
	Ducklings    []Duckling
}

var (
	App         AppConfiguration
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

	if LoadProject() == false {
		return false
	}

	LoadDuckfiles()

	return true
}

func LoadProject() bool {
	dir, _ := os.Getwd()

	if !ExistsConfIn(dir) {
		fmt.Println("No duck repo found in", dir)
		return false
	}

	LoadAppConfig()

	LoadProjectConfig(dir)
	return true
}

func (p *Configuration) Write() {
	var root string
	if len(Conf.ProjectRoot) == 0 {
		root, _ = os.Getwd()
		fmt.Println("Taking", root, "as project path")
	} else {
		root = Conf.ProjectRoot
	}
	b, _ := json.Marshal(*p)
	ioutil.WriteFile(root+"/.duck/project.conf", b, 0644)
}

func LoadAppConfig() {
	LoadFileJson("/etc/duck.conf", &App)
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
			str := BLUE + " - " + tmp.Label + END_STYLE
			verbose(str)
			Ducklings = append(Ducklings, tmp)
		}
	}
}

func InstallDuckling(args ...string) {
	//params
	if !LoadProject() {
		return
	}
	globalMode := false
	forceMode := false
	var ducklings []string
	var errors []string
	for _, arg := range args {
		if arg[0] != '-' {
			ducklings = append(ducklings, arg)
			continue
		}

		switch arg {
		case "-g":
			globalMode = true
			fmt.Println("installing globally:", globalMode)
			break
		case "-f":
			forceMode = true
			break
		default:
			fmt.Println(RED+"Unknown parameter", BLUE+arg+END_STYLE)
			break
		}
	}

	if len(ducklings) == 0 {
		if len(Conf.Ducklings) > 0 {
			for _, duckling := range Conf.Ducklings {
				args = append(args, duckling)
			}
			InstallDuckling(args...)
		} else {
			fmt.Println(RED + "No ducklings to install" + END_STYLE)
		}
		return
	}

	//print enabled modes
	if forceMode == true {
		fmt.Println(BLUE+"-f"+END_STYLE, ": force dependencies to install", GREEN+"(active)"+END_STYLE)
	}

	//for each requested duckling
	for _, arg := range ducklings {
		fmt.Print("\rinstall ", BLUE+arg+END_STYLE, "...")
		installed := false

		cmd := exec.Command("mkdir", Conf.ProjectRoot+"/.ducklings")
		cmd.Run()
		path := arg + ".duckling"
		cmd = exec.Command("mkdir", Conf.ProjectRoot+"/.ducklings/"+strings.Split(path, "/")[0])
		cmd.Run()

		for i, repo := range App.Repos {
			fmt.Print("\rinstall ", BLUE+arg+END_STYLE, "... (", i+1, "/", len(App.Repos), ")")
			filePath := Conf.ProjectRoot + "/.ducklings/" + path

			url := repo.URL + path + "?" + fmt.Sprintf("%d", rand.Int())
			//fmt.Println("Fetching", url)
			//
			cmd = exec.Command("curl", url, "-o", filePath, "-f")
			err := cmd.Run()
			if err != nil {
				continue
			} else {
				var tmp Duckfile
				LoadFileJson(filePath, &tmp)
				for _, dep := range tmp.Dependencies {
					resolved := false
					if forceMode == false {
						//check if the dependency is already satisfied
						for _, available := range Conf.Ducklings {
							if available == dep {
								resolved = true
								break
							}
						}
					}
					if resolved == false {
						fmt.Println("\r"+YELLOW+"installing dependency", BLUE+dep+END_STYLE)
						InstallDuckling(dep)
					}
				}
				installed = true
				fmt.Println("\r"+GREEN+"installed", BLUE+arg, GREEN+"from", YELLOW+repo.Name+END_STYLE)
				break
			}
		}
		if installed == false {
			msg := RED + "Not found " + BLUE + arg + RED + " in any repository." + END_STYLE
			errors = append(errors, msg)
			continue
		}

		found := false
		for _, active := range Conf.Ducklings {
			if arg == active {
				found = true
				break
			}
		}
		if found != true {
			Conf.Ducklings = append(Conf.Ducklings, arg)
			Conf.Write()
		}
	}
	for _, msg := range errors {
		fmt.Println(msg)
	}
}

func UninstallDuckling(args ...string) bool {
	if !LoadProject() {
		return false
	}

	var newDucklings []string

	for _, duckling := range Conf.Ducklings {
		found := false
		for _, arg := range args {
			if duckling == arg {
				fmt.Println(YELLOW+"Deleting", BLUE+duckling+END_STYLE)
				found = true
				break
			}
		}
		if found == false {
			newDucklings = append(newDucklings, duckling)
		}
	}

	Conf.Ducklings = newDucklings
	Conf.Write()
	return true
}

func PrintRepos() {
	Init()
	for _, repo := range App.Repos {
		fmt.Println(BLUE+repo.Name+END_STYLE, YELLOW+repo.URL+END_STYLE)
	}
	fmt.Println("total:", len(App.Repos))
}

func PrintDucklings() {
	Init()
	for _, duckling := range Conf.Ducklings {
		fmt.Println(duckling)
	}
	fmt.Println("total:", len(Conf.Ducklings))
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
	DUCK_DIR := ".duck"
	DUCK_PERM := os.FileMode(0740)

	dir, _ := os.Getwd()

	if ExistsConfIn(dir) {
		fmt.Println("error: this directory already has a duck repo.")
		return
	}

	fmt.Println("initializing a new repo in", dir)

	err := os.Mkdir(DUCK_DIR, DUCK_PERM)

	checkErr(err)

	NewConf.ProjectRoot = dir
	NewConf.Name = askProperty("Name: ")
	NewConf.Lang = askProperty("Lang: ")
	NewConf.VersionMajor = "1.0"
	NewConf.VersionMinor = "0"
	NewConf.Env = "dev"
	NewConf.Ducklings = []string{}
	NewConf.Main = askProperty("Main: ")

	NewConf.Write()
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
