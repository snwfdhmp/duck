package projects

import (
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/spf13/afero"
)

var (
	storageDir = ".duck"
	confName   = "conf.ini"

	HasErrors bool = false

	IsInside     bool
	Root         string
	Storage      string
	Config       *ini.File
	ConfigPath   string
	Packages     []string
	PackagesDir  string = "packages"
	PackagesPath string
)

func init() {
	var err error
	IsInside, Root, err = findProject()
	if err != nil {
		HasErrors = true
	}

	Storage = makeStoragePath(Root)
	PackagesPath = filepath.Join(Storage, PackagesDir)
	ConfigPath = makeConfPath(Root)
	Config, err = ini.Load(ConfigPath)
	if err != nil {
		HasErrors = true
	}
}

func SaveConfig() error {
	return Config.SaveTo(ConfigPath)
}

func makeStoragePath(s string) string {
	return filepath.Join(s, storageDir)
}

func makeConfPath(s string) string {
	return filepath.Join(makeStoragePath(s), confName)
}

func findProject() (bool, string, error) {
	root, err := filepath.Abs("./")
	if err != nil {
		return false, "", err
	}

	exists, err := afero.Exists(fs, makeConfPath(root))
	if err != nil {
		return false, "", err
	}

	for !exists && root != "/" {
		filepath.Dir(root)
		exists, err = afero.Exists(fs, makeConfPath(root))
		if err != nil {
			return false, "", err
		}
	}

	return exists, root, nil
}

func IsHealthy(path string) (bool, error) {
	fs := afero.NewOsFs()

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}

	toTest := []string{
		absolutePath + "/.duck",
		absolutePath + "/.duck/conf.ini",
		absolutePath + "/.duck/packages",
	}

	for _, s := range toTest {
		exists, err := afero.Exists(fs, s)
		if err != nil {
			return false, err
		} else if !exists {
			return false, nil
		}
	}

	config, err := ini.Load(absolutePath + "/.duck/conf.ini")
	if err != nil {
		return false, nil
	}

	projectSection, err := config.GetSection("project")
	if err != nil {
		return false, nil
	}

	_, err = projectSection.GetKey("name")
	if err != nil {
		return false, nil
	}

	pathKey, err := projectSection.GetKey("path")
	if err != nil || pathKey.Value() != absolutePath {
		return false, nil
	}

	return true, nil
}
