package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"runtime"
	"strings"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		fmt.Println(lookCname(checkDomain(scanner.Text())))
	}

}

func checkDomain(s string) string {
	re := regexp.MustCompile("^https://|^http://|/|^'*'.")
	replace := re.ReplaceAllString(s, "")
	return replace

}

func lookCname(d string) string {
	cName, _ := net.LookupCNAME(d)
	dot := strings.TrimRight(cName, ".")
	if dot != d {
		return Yellow + d + Reset + "  Cname:  " + Green + dot + Reset
	}
	return Yellow + d + Reset + "  Cname:  " + Red + dot + Reset
}

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}
