package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/snwfdhmp/duck/pkg/data"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const (
	DefaultRepoName = "core"
	DefaultRepoURL  = "http://raw.githubusercontent.com/snwfdhmp/duck-core/master/"
)

var (
	fs = afero.NewOsFs()

	packageExtension = ".duckpkg.ini"
	err              error
)

func File(target, pkg string) (afero.File, error) {
	arr := strings.Split(pkg, "/")

	packageName := arr[len(arr)-1]

	paths := append([]string{target}, arr[:len(arr)-1]...)
	authorFolder := filepath.Join(paths...)

	exists, err := afero.Exists(fs, authorFolder)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = fs.MkdirAll(authorFolder, 0740)
		if err != nil {
			return nil, err
		}
	}

	packagePath := filepath.Join(authorFolder, packageName+packageExtension)

	exists, err = afero.Exists(fs, packagePath)
	if err != nil {
		return nil, err
	}

	var openFile func(string) (afero.File, error)
	if !exists {
		openFile = fs.Create
	} else {
		openFile = fs.Open
	}

	return openFile(packagePath)
}

func DownloadMany(target string, pkgs []string) []bool {
	var err []bool

	for _, pkg := range pkgs {
		err = append(err, Download(target, pkg))
	}

	return err
}

func Download(target, pkg string) bool {
	installed := false

	if len(data.Repos) == 0 {
		fmt.Println("No repository configured. Installing default repository...")
		_, err = data.AddRepo(DefaultRepoName, DefaultRepoURL)
		if err != nil {
			color.Red("Could not add default repo : " + err.Error())
			return false
		}
	}
	for _, url := range data.Repos {
		body, err := downloadFrom(pkg, url)
		if err != nil {
			color.Red("Could not download from " + url)
			continue
		}

		out, err := File(target, pkg)
		if err != nil {
			color.Red("Could not create file " + err.Error())
			continue
		}

		_, err = io.Copy(out, body)
		if err != nil {
			color.Red("Could not copy data " + err.Error())
			continue
		}
		installed = true
	}
	return installed
}

func Load(target, pkg string) (string, *ini.File, error) {
	_, err := File(target, pkg)
	if err != nil {
		return "", nil, err
	}
	path := filepath.Join(target, pkg, packageExtension)
	file, err := ini.Load(path)
	return path, file, err
}

func Create(target, pkg string, command map[string]string) error {
	pkgPath, pkgFile, err := Load(target, pkg)
	if err != nil {
		return err
	}

	createCommand(pkgFile, command)

	return pkgFile.SaveTo(pkgPath)
}

func List(dir string) []string {
	var paths []string
	fi, err := afero.ReadDir(fs, dir)
	if err != nil {
		return []string{}
	}

	for i := 0; i < len(fi); i++ {
		path := filepath.Join(dir, fi[i].Name())
		if fi[i].IsDir() {
			paths = append(paths, List(path)...)
			continue
		}
		if path[len(path)-len(packageExtension):] == packageExtension {
			paths = append(paths, path)
		}
	}
	return paths
}

type Command struct {
	Name      string
	Shortcut  string
	Cmd       string
	ShortHelp string
	LongHelp  string
}

func ReadDirs(dirs []string) ([]Command, error) {
	var paths []string
	for _, dir := range dirs {
		paths = append(paths, List(dir)...)
	}
	return ReadMany(paths)
}

func ReadMany(paths []string) ([]Command, error) {
	var cmds []Command
	for i := 0; i < len(paths); i++ {
		cfg, err := ini.Load(paths[i])
		if err != nil {
			color.Red(err.Error())
		}
		sections := cfg.Sections()
		for j := 0; j < len(sections); j++ {
			if sections[j].Name() == "DEFAULT" {
				continue
			}
			alreadyExists := false
			for _, tmp := range cmds {
				if tmp.Name == sections[j].Name() {
					alreadyExists = true
					break
				}
			}
			if alreadyExists {
				continue
			}
			cmds = append(cmds, Command{
				Name:      sections[j].Name(),
				Cmd:       sections[j].Key("cmd").String(),
				Shortcut:  sections[j].Key("shortcut").String(),
				ShortHelp: sections[j].Key("help").String(),
				LongHelp:  sections[j].Key("longHelp").String(),
			})
		}
	}
	return cmds, nil
}

func CreateCobraCommands(root *cobra.Command, cmds []Command) {
	for i := 0; i < len(cmds); i++ {
		var tmpCmd = &cobra.Command{
			Use:     cmds[i].Name,
			Short:   cmds[i].ShortHelp,
			Long:    cmds[i].LongHelp,
			Aliases: []string{cmds[i].Shortcut},
			Run: func(cmd *cobra.Command, args []string) {
				i := cmd.DuckCmdIndex
				commonArgs := []string{"-c", cmds[i].Cmd, cmd.Use}
				args = append(commonArgs, args...)
				shell := os.Getenv("SHELL")
				execCmd := exec.Command(shell, args...)
				execCmd.Stderr = os.Stderr
				execCmd.Stdout = os.Stdout
				execCmd.Stdin = os.Stdin
				err := execCmd.Run()
				if err != nil {
					color.Red("An error occured while executing this command. (" + shell + ") Error: " + err.Error())
					os.Exit(1)
				}
			},
			DuckCmdIndex: i,
		}
		root.AddCommand(tmpCmd)
	}
}

func CreateMany(target, pkg string, commands []map[string]string) error {
	pkgPath, pkgFile, err := Load(target, pkg)
	if err != nil {
		return err
	}

	for _, command := range commands {
		err = createCommand(pkgFile, command)
		if err != nil {
			return err
		}
	}

	return pkgFile.SaveTo(pkgPath)
}

func createCommand(file *ini.File, command map[string]string) error {
	section, err := file.NewSection(command["name"])
	if err != nil {
		return err
	}
	for key, value := range map[string]string{
		"cmd":      command["cmd"],
		"shortcut": command["shortcut"],
		"help":     command["help"],
	} {
		_, err = section.NewKey(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadFrom(pkg, url string) (io.ReadCloser, error) {
	packageURL := url + pkg + packageExtension
	resp, err := http.Get(packageURL) //test errors
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Body, nil
}
