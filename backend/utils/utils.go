package utils

import (
	"fmt"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func WarningHandling(types string, msg string) {
	fmt.Println(Yellow + "[" + types + "] " + msg)
	fmt.Print(Reset)

}

func ErrorHandling(err error) {
	fmt.Println(Red + err.Error())
	fmt.Print(Reset)
}
func PrintHandling(types string, msg string) {
	fmt.Println(Green + "[" + types + "] " + msg)
	fmt.Print(Reset)
}
