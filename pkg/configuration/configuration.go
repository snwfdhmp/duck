package conf

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/snwfdhmp/duck/pkg/logger"
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
	log         logger.Logger
)

//Run execute the "duck config {command}"
//command. It just prints config values
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

//Init Load project and packages
func Init() error {

	if LoadProject() == false {
		return errors.New("cannot load project and/or app configuration")
	}

	LoadPackages()

	return nil
}

//LoadProject Loads App Config and Project Config
//into App and Conf
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

//Writer is the Writer for Conf file
func (p *Configuration) Write() {
	var root string
	if len(Conf.ProjectRoot) == 0 {
		root, err := os.Getwd()
		log.Err(err, "Cannot get current directory")

		fmt.Println("Taking", root, "as project path")
	} else {
		root = Conf.ProjectRoot
	}

	//write file and check err
	b, err := json.Marshal(*p)
	log.Fatal(err, "Cannot write config file")

	ioutil.WriteFile(root+"/.duck/project.conf", b, 0644)
}

//Write writes AppConfiguration to /etc/duck/duck.conf
func (a *AppConfiguration) Write() {
	b, err := json.Marshal(*a)
	log.Fatal(err, "Cannot write app configuration file")

	ioutil.WriteFile("/etc/duck.conf", b, 0644)
}

//LoadAppConfig loads duck configuration file
//from /etc/duck.conf and puts it into App
func LoadAppConfig() {
	LoadFileJson("/etc/duck.conf", &App)
}

//LoadFileJson loads a JSON file into its correctly typed interface
func LoadFileJson(path string, object interface{}) error {
	//open file
	file, err := os.Open(path)

	//if it doesn't exist, error
	if os.IsNotExist(err) {
		fmt.Println("Error: not found", path, "found")
		return errors.New("not found")
	}
	if log.Err(err, "Cannot load file "+path) {
		return errors.New("cannot load")
	}

	//decode it
	decoder := json.NewDecoder(file)

	//parse it into object
	err = decoder.Decode(object)
	if log.Err(err, "Cannot decode "+path) {
		return errors.New("cannot decode")
	}

	return nil
}

// LoadProjectConfig loads project configuration (./.duck/project.conf) file into Conf
func LoadProjectConfig(dir string) error {
	path := dir + "/.duck/project.conf"
	err := LoadFileJson(path, &Conf)
	log.Check(err)
	//fmt.Println("Loaded conf for project", Conf.Name, "in language", Conf.Lang)
	return nil
}

//BuildPkgPath returns the normalized path for a package named str
func BuildPkgPath(str string) string {
	return Conf.ProjectRoot + "/" + Conf.PkgLocation + "/" + str + ".pkg"
}

//verbose prints str if verboseMode is active (== true)
func verbose(str string) {
	if verboseMode {
		fmt.Println(str)
	}
}

//LoadPackages loads project.conf packages
func LoadPackages() {
	//create array that will store every Ling for the project
	Lings = []Ling{}

	//for each packages defined in project conf
	for _, pkg := range Conf.Packages {
		//generate the pkg path ($path/$pkglocation/pkg)
		path := BuildPkgPath(pkg)

		//if we're in verbose mode
		//	print infos about current pkg
		verbose(logger.YELLOW + "from " + pkg + logger.END_STYLE)

		var pkg Pkg              //create a pkg object
		LoadFileJson(path, &pkg) //load the pkg file into it

		//foreach lings in it, add it to our Lings array
		for _, ling := range pkg.Lings {
			verbose(logger.BLUE + " - " + ling.Label + logger.END_STYLE) // if verbose, print ling
			Lings = append(Lings, ling)                                  //append ling to Lings array
		}
	}
}

//isInstalled returns whether a package named pkg
//is currently installed
func isInstalled(pkg string) bool {
	for _, tmp := range Conf.Packages { //if dep is in Conf.Packages
		if tmp == pkg {
			return true //exit for loop
		}
	}
	return false
}

//InstallPkg installs a package
//forceMode force to reinstall dependencies
func InstallPkg(pkg string, forceMode bool) {
	fmt.Print("\rinstalling ", logger.BLUE+pkg+logger.END_STYLE, "...")
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
		fmt.Print("\rinstalling ", logger.BLUE+pkg+logger.END_STYLE, "... (", i+1, "/", len(App.Repos), ")")
		filePath := BuildPkgPath(pkg)

		url := repo.URL + path + "?" + fmt.Sprintf("%d", rand.Int())
		//fmt.Println("Fetching", url)

		//download from url to filePath with "-f" enabling exit error when 404
		cmd = exec.Command("curl", url, "-o", filePath, "-f")
		err := cmd.Run()

		//if curl failed, try the next repo
		if err != nil {
			continue
		}

		//Load file into Pkg object
		var tmp Pkg
		LoadFileJson(filePath, &tmp)

		for _, dep := range tmp.Dependencies {
			if forceMode == false || !isInstalled(dep) {
				//check if the dependency is already satisfied
				fmt.Println("\r"+logger.YELLOW+"installing dependency", logger.BLUE+dep+logger.END_STYLE)
				InstallPkg(dep, false)
			}
		}

		installed = true
		fmt.Println("\r"+logger.GREEN+"installed", logger.BLUE+pkg, logger.GREEN+"from", logger.YELLOW+repo.Name+logger.END_STYLE)
		break
	}
	if installed == false {
		fmt.Println(logger.RED + "\rnot found " + logger.BLUE + pkg + logger.RED + " in any repository." + logger.END_STYLE)
		return
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

//InstallPkgs install every pkgs
//This function also receives options
//Options:
//	-f	force to reinstall dependencies
//	-g	(incomming) install globally
func InstallPkgs(args ...string) {
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
			fmt.Println(logger.RED+"Unknown parameter", logger.BLUE+arg+logger.END_STYLE)
			break
		}
	}

	if len(pkgs) == 0 { //if no pkgs in command
		if len(Conf.Packages) > 0 { // reinstall project pkgs
			//add all pkgs to args list
			for _, pkg := range Conf.Packages {
				args = append(args, pkg)
			}
			InstallPkgs(args...)
		} else { //or print error if project hasn't pkgs
			fmt.Println(logger.RED + "No ducklings to install" + logger.END_STYLE)
		}
		return
	}

	//print enabled modes
	if forceMode == true {
		fmt.Println(logger.BLUE+"-f"+logger.END_STYLE, ": force dependencies to install", logger.GREEN+"(active)"+logger.END_STYLE)
	}

	//for each requested pkg
	for _, pkg := range pkgs {
		InstallPkg(pkg, forceMode)
	}

	//print every error encountered
	for _, msg := range errors {
		fmt.Println(msg)
	}
}

//UninstallPkgs removes some .duckpkg from the project.conf
//args is the List of the .duckpkg to uninstall
func UninstallPkgs(args ...string) {
	if !LoadProject() {
		log.FatalString("Cannot load project")
	}

	//future list of installed packages
	var newPackages []string

	//foreach installed package
	for _, pkg := range Conf.Packages {
		found := false
		//search for it into command args
		for _, arg := range args {
			if pkg == arg { //if found don't add it to future list
				fmt.Println(logger.GREEN+"deleted", logger.RED+pkg+logger.END_STYLE)
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
}

//PrintRepos prints every repo configured in /etc/duck.conf
func PrintRepos() {
	Init()

	//print each repo in App.Repos
	for _, repo := range App.Repos {
		fmt.Println(logger.BLUE+repo.Name+logger.END_STYLE, logger.YELLOW+repo.URL+logger.END_STYLE)
	}
	fmt.Println("total:", len(App.Repos))
}

//PrintPackagegs returns the list of
//every package in Conf.Packages
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

//GetCommands returns the array of commands
//set up for label by looking in Lings
func GetCommands(label string) []string {
	//looking for the commands corresponding to the label
	for _, ling := range Lings {
		if ling.Label == label { // if scheme's label is input, return it
			return ling.Commands
		} else { //else look in its aliases
			for _, alias := range ling.Aliases {
				if alias == label {
					return ling.Commands
				}
			}
		}
	}

	//if nothing found, return an error
	return []string{"echo " + logger.RED + "Unknown command '" + label + "'" + logger.END_STYLE} //@todo handle errors better
}

//ExistsConfIn returns whether there's a .duck
//in the current directory
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

//AskConf asks the user for configuration
//values to init or modify a project conf
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

	log.Fatal(err, "Cannot create "+DUCK_DIR+" in "+dir)

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

/**
 * Add a new repo to App conf
 * @param {string} repo  The repo
 */
func AddRepo(name string, url string) {
	//Load project repos
	LoadProject()

	//push new repo
	App.Repos = append(App.Repos, Repo{Name: name, URL: url})

	//write changes
	App.Write()
}

func getRidNewLine(str string) string {
	return str[:len(str)-1]
}

func askProperty(prompt string) string {
	reader := bufio.NewReader(os.Stdin) //reader initialized for stdin
	fmt.Print(prompt)
	tmp, err := reader.ReadString('\n')
	log.Fatal(err, "Cannot read")

	return getRidNewLine(tmp)
}
