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
	Packages     []string
	Name         string
	Lang         string
	VersionMajor string
	VersionMinor string
	Env          string
	Main         string
	PkgLocation  string
}

//@todo rename Schemes to Ducklings ("caneton" in French (duck's children))
type Ling struct {
	Label       string
	Commands    []string
	Description string
	Aliases     []string
}

type Pkg struct {
	Dependencies []string
	Lings        []Ling
}

var (
	App         AppConfiguration
	Conf        Configuration
	Lings       []Ling
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

	LoadPackages()

	return true
}

/**
 * Loads App Config and Project Config
 *  into App and Conf
 */
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

/**
 * @brief      Writer for Conf file
 *
 */
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

/**
 * Load App Config from /etc/duck.conf
 *  into App
 */
func LoadAppConfig() {
	LoadFileJson("/etc/duck.conf", &App)
}

//load a JSON file into its correctly typed interface
func LoadFileJson(path string, object interface{}) bool {
	//open file
	file, err := os.Open(path)

	//if it doesn't exist, error
	if os.IsNotExist(err) {
		fmt.Println("Error: not found", path, "found")
		return false
	}
	checkErr(err)

	//decode it
	decoder := json.NewDecoder(file)

	//parse it into object
	err = decoder.Decode(object)
	checkErr(err)
	return true
}

/**
 * load Project Configuration (./.duck/project.conf) file into Conf
 */
func LoadProjectConfig(dir string) bool {
	path := dir + "/.duck/project.conf"
	LoadFileJson(path, &Conf)
	//fmt.Println("Loaded conf for project", Conf.Name, "in language", Conf.Lang)
	return true
}

func BuildPkgPath(str string) string {
	return Conf.ProjectRoot + "/" + Conf.PkgLocation + "/" + str + ".pkg"
}

func verbose(str string) {
	if verboseMode {
		fmt.Println(str)
	}
}

func LoadPackages() {
	//create array that will store every Ling for the project
	Lings = []Ling{}

	//for each packages defined in project conf
	for _, pkg := range Conf.Packages {
		//generate the pkg path ($path/$pkglocation/pkg)
		path := BuildPkgPath(pkg)

		//if we're in verbose mode
		//	print infos about current pkg
		verbose(YELLOW + "from " + pkg + END_STYLE)

		var pkg Pkg              //create a pkg object
		LoadFileJson(path, &pkg) //load the pkg file into it

		//foreach lings in it, add it to our Lings array
		for _, ling := range pkg.Lings {
			verbose(BLUE + " - " + ling.Label + END_STYLE) // if verbose, print ling
			Lings = append(Lings, ling)                    //append ling to Lings array
		}
	}
}

func InstallPkg(args ...string) {
	//stop if no duck conf or project conf
	if !LoadProject() {
		return
	}

	//init modes
	globalMode := false
	forceMode := false

	var pkgs []string   //pkgs asked to install will be stored in this array
	var errors []string //errors catching

	//foreach argument
	for _, arg := range args {
		//if it's not an option, treat it like a pkg
		if arg[0] != '-' {
			pkgs = append(pkgs, arg)
			continue
		}

		//handle options
		switch arg {
		case "-g": //not doing anything yet
			globalMode = true
			fmt.Println("installing globally:", globalMode)
			break
		case "-f": //will force reinstall dependencies for every pkg
			forceMode = true
			break
		default:
			fmt.Println(RED+"Unknown parameter", BLUE+arg+END_STYLE)
			break
		}
	}

	if len(pkgs) == 0 { //if no pkgs in command
		if len(Conf.Packages) > 0 { // reinstall project pkgs
			//add all pkgs to args list
			for _, pkg := range Conf.Packages {
				args = append(args, pkg)
			}
			InstallPkg(args...)
		} else { //or print error if project hasn't pkgs
			fmt.Println(RED + "No ducklings to install" + END_STYLE)
		}
		return
	}

	//print enabled modes
	if forceMode == true {
		fmt.Println(BLUE+"-f"+END_STYLE, ": force dependencies to install", GREEN+"(active)"+END_STYLE)
	}

	//for each requested pkg
	for _, pkg := range pkgs {
		fmt.Print("\rinstalling ", BLUE+pkg+END_STYLE, "...")
		installed := false

		//create PkgLocation directory
		cmd := exec.Command("mkdir", Conf.ProjectRoot+"/"+Conf.PkgLocation)
		cmd.Run()

		//build file path for pkg (ex: snwfdhmp/go => snwfdhmp/go.pkg)
		path := pkg + ".pkg"

		//create pkg @author directory
		cmd = exec.Command("mkdir", Conf.ProjectRoot+"/"+Conf.PkgLocation+"/"+strings.Split(path, "/")[0])
		cmd.Run()

		//foreach repo in /etc/duck.conf
		for i, repo := range App.Repos {
			fmt.Print("\rinstalling ", BLUE+pkg+END_STYLE, "... (", i+1, "/", len(App.Repos), ")")
			filePath := BuildPkgPath(pkg)

			url := repo.URL + path + "?" + fmt.Sprintf("%d", rand.Int())
			//fmt.Println("Fetching", url)

			//download from url to filePath with "-f" enabling exit error when 404
			cmd = exec.Command("curl", url, "-o", filePath, "-f")
			err := cmd.Run()

			//if no
			if err != nil {
				continue
			}

			//Load file into Pkg object
			var tmp Pkg
			LoadFileJson(filePath, &tmp)

			for _, dep := range tmp.Dependencies {
				resolved := false
				if forceMode == false {
					//check if the dependency is already satisfied
					for _, available := range Conf.Packages { //if dep is in Conf.Packages
						if available == dep {
							resolved = true //exit for loop
							break
						}
					}
				}
				if resolved == false {
					fmt.Println("\r"+YELLOW+"installing dependency", BLUE+dep+END_STYLE)
					InstallPkg(dep)
				}
			}

			installed = true
			fmt.Println("\r"+GREEN+"installed", BLUE+pkg, GREEN+"from", YELLOW+repo.Name+END_STYLE)
			break
		}
		if installed == false {
			msg := RED + "\rNot found " + BLUE + pkg + RED + " in any repository." + END_STYLE
			errors = append(errors, msg)
			continue
		}

		//if the pkg was already in Conf.Packages, exit
		found := false
		for _, active := range Conf.Packages {
			if pkg == active {
				found = true
				break
			}
		}

		//else push it into Conf.Packages & write changes
		if found == false {
			Conf.Packages = append(Conf.Packages, pkg)
			Conf.Write()
		}
	}

	//print every error encountered
	for _, msg := range errors {
		fmt.Println(msg)
	}
}

/**
 * @brief      Remove some .duckpkg from the project.conf
 *
 * @param      args  List of the .duckpkg to uninstall
 *
 * @return     { exitSuccess }
 */
func UninstallDuckling(args ...string) bool {
	if !LoadProject() {
		return false
	}

	//future list of installed packages
	var newPackages []string

	//foreach installed package
	for _, pkg := range Conf.Packages {
		found := false
		//search for it into command args
		for _, arg := range args {
			if pkg == arg { //if found don't add it to future list
				fmt.Println(YELLOW+"Deleting", BLUE+pkg+END_STYLE)
				found = true
				break
			}
		}
		if found == false { //if not found this means pkg must be kept installed
			newPackages = append(newPackages, pkg)
		}
	}

	Conf.Packages = newPackages //change Packages to future list
	Conf.Write()                //write changes
	return true
}

/**
 * @brief      Print every repo configured in /etc/duck.conf
 */
func PrintRepos() {
	Init()

	//print each repo in App.Repos
	for _, repo := range App.Repos {
		fmt.Println(BLUE+repo.Name+END_STYLE, YELLOW+repo.URL+END_STYLE)
	}
	fmt.Println("total:", len(App.Repos))
}

func PrintPackages() {
	Init()
	for _, pkg := range Conf.Packages {
		fmt.Println(pkg)
	}
	fmt.Println("total:", len(Conf.Packages))
	verboseMode = true
	LoadPackages()
	verboseMode = false
}

func GetCommands(label string) []string {
	//looking for the commands corresponding to the label
	for _, val := range Lings {
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
	NewConf.PkgLocation = ".duck/pkg"
	NewConf.Packages = []string{}
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
