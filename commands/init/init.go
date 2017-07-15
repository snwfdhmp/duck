package InitCmd

import (
	"../../configuration"
	"fmt"
	"os"
)

func Run() bool {
	DUCK_DIR := ".duck"
	DUCK_PERM := os.FileMode(0740)

	dir, _ := os.Getwd()

	if conf.ExistsConfIn(dir) {
		fmt.Println("error: this directory already has a duck repo.")
		return false
	}

	fmt.Println("initializing a new repo in", dir)
	err := os.Mkdir(DUCK_DIR, DUCK_PERM)
	checkErr(err)

	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
