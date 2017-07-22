package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
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

type Logger struct{}

//Check prints "error: {err}" if err != nil
//It returns err != nil
func (Logger) Check(err error) bool {
	if err == nil {
		return false
	}
	fmt.Println(RED+"error:", err, END_STYLE)
	return true
}

//Err prints msg to stdout if err != nil
//and logs err, it returns err != nil
func (Logger) Err(err error, msg string) bool {
	if err == nil {
		return false
	}

	fmt.Println(RED + msg + END_STYLE) //print nice error to user
	log.Println(err)                   //log error
	return true
}

//Fatal prints msg to stdout if err != nil, logs
//err and exit with exit 1.
func (Logger) Fatal(err error, msg string) {
	if err != nil {
		fmt.Println(RED + "Fatal: " + msg + END_STYLE)
		log.Fatalln(err)
	}
}

//FatalString calls l.Fatal with msg as err
func (l Logger) FatalString(msg string) {
	err := errors.New(msg)
	l.Fatal(err, msg)
}
