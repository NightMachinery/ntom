/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	//"bytes"
	. "fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	//"bytes"
	"os"
	// "github.com/logrusorgru/aurora"
	//"github.com/bitfield/script"
	//"os/exec"
)

func main() {
	topS := color.NewRGBStyle(color.RGB(0, 0, 0), color.RGB(200, 255, 200))
	matchS := color.NewRGBStyle(color.RGB(255, 120, 0), color.RGB(255, 255, 255))
	filePath := os.Args[1]
	ln, _ := strconv.Atoi(os.Args[2])
	ln -= 1 // Make it zero-based
	match := os.Args[3]
	relDir := ""
	if len(os.Args) >= 5 {
		relDir = os.Args[4] + "/"
	}
	prevLines := 6
	afterLines := 15
	//Printf("ln %v\nmatch %v\nfile %v\n", ln, match, filePath)

	if _, err := os.Stat(filePath); err != nil {
		if _, err := os.Stat(relDir + filePath); err == nil {
			filePath = relDir + filePath
		} else {
			log.Fatalln("ntom: File supplied did not exist.")
		}
	}
	Println(filePath)
	topS.Println(match)
	Println()
	//Println() // The color would leak with Printf
	var fileBytes, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fileLines := strings.Split(string(fileBytes), "\n")
	if ln != 1 {
		ps := ln - prevLines
		if ps < 1 {
			ps = 1
		}
		Println(strings.Join(fileLines[ps:ln], "\n"))
	}
	matchS.Println(fileLines[ln])
	afterLinesTLen := len(fileLines) - 1
	if afterLinesTLen > ln {
		ae := ln + afterLines
		if ae > afterLinesTLen {
			ae = afterLinesTLen
		}
		Println(strings.Join(fileLines[(ln+1):(ae+1)], "\n"))
	}

	os.Exit(0)
}
