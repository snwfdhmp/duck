package InitCmd

import (
	"fmt"
	"os"
	"../../configuration"
)

func Run() bool {
	DUCK_DIR := ".duck"
	DUCK_PERM := os.FileMode(0740)

	dir, _ := os.Getwd()

	if(conf.ExistsConfIn(dir)) {
		fmt.Println("error: this directory already has a duck repo.")
		return false
	}

	fmt.Println("initializing a new repo in", dir)
	err := os.Mkdir(DUCK_DIR, DUCK_PERM)
	checkErr(err)

	fmt.Println("downloading ")
	return true
}

func checkErr(err error) {
	if (err != nil) {
		panic(err)
	}
}