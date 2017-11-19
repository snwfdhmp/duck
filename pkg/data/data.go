package data

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
)

var (
	fs  = afero.NewOsFs()
	err error

	HasErrors bool = false

	//config
	storageDir   string = ".duck"
	dataFileName string = "data.ini"
	packagesDir  string = "packages"

	//accesseurs
	StoragePath  string
	File         *ini.File
	Path         string
	Projects     *ini.Section
	PackagesPath string
	Repos        map[string]string
)

func getStoragePath() (string, error) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := filepath.Join(home, storageDir)

	exists, err := afero.Exists(fs, path)
	if err != nil {
		return "", err
	}
	if exists {
		return path, nil
	}

	err = fs.MkdirAll(path, 0755)
	if err != nil {
		return "", err
	}

	return path, nil
}

func loadFile(path string) (*ini.File, error) {
	exists, err := afero.Exists(fs, path)
	if err != nil {
		return &ini.File{}, err
	}
	if !exists {
		_, err := fs.Create(path)
		if err != nil {
			return &ini.File{}, err
		}
	}

	config, err := ini.Load(path)
	if err != nil {
		return &ini.File{}, err
	}

	_, err = config.GetSection("projects")
	if err != nil {
		Projects, err = config.NewSection("projects")
		if err != nil {
			return &ini.File{}, err
		}
		err := config.SaveTo(path)
		if err != nil {
			return &ini.File{}, err
		}
	}

	return config, nil
}

func getRepos() (map[string]string, error) {
	reposSection, err := File.GetSection("repos")
	if err != nil {
		reposSection, err = File.NewSection("repos")
		if err != nil {
			return nil, err
		}
	}

	repos := make(map[string]string)

	for _, repo := range reposSection.Keys() {
		repos[repo.Name()] = repo.Value()
	}

	return repos, nil
}

func AddRepo(name, url string) (*ini.Key, error) {
	reposSection, err := File.GetSection("repos")
	if err != nil {
		reposSection, err = File.NewSection("repos")
		if err != nil {
			return nil, err
		}
	}

	key, err := reposSection.NewKey(name, url)
	if err != nil {
		return nil, err
	}
	err = Save()
	if err != nil {
		return nil, err
	}

	LoadRepos()

	return key, Save()
}

func Save() error {
	return File.SaveTo(Path)
}

func LoadPackages() {
	//Create PackagesPath if not exists
	exists, err := afero.Exists(fs, PackagesPath)
	if err != nil {
		HasErrors = true
	}
	if !exists {
		err := fs.MkdirAll(PackagesPath, 0740)
		if err != nil {
			HasErrors = true
		}
	}
}

func LoadRepos() {
	Repos, err = getRepos()
	if err != nil {
		HasErrors = true
	}
}

func LoadFile() {
	File, err = loadFile(Path)
	if err != nil {
		HasErrors = true
	}
}

func init() {
	StoragePath, err = getStoragePath()
	if err != nil {
		HasErrors = true
	}

	Path = filepath.Join(StoragePath, dataFileName)
	PackagesPath = filepath.Join(StoragePath, packagesDir)

	LoadFile()
	LoadPackages()
	LoadRepos()
}
