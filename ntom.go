/// 2>/dev/null ; exec gorun "$0" "$@"

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
	l := log.New(os.Stderr, "", 0)
	topS := color.NewRGBStyle(color.RGB(0, 0, 0), color.RGB(200, 255, 200))
	matchS := color.NewRGBStyle(color.RGB(255, 120, 0), color.RGB(255, 255, 255))
	filePath := os.Args[1]
	ln, err := strconv.Atoi(os.Args[2])
	if err != nil {
		l.Fatal(err)
	}
	ln -= 1 // Make it zero-based
	match := os.Args[3]
	relDir := ""
	if len(os.Args) >= 5 {
		relDir = os.Args[4] + "/"
	}
	mode := 0
	if len(os.Args) >= 6 {
		_mode, err := strconv.Atoi(os.Args[5])
		if err != nil {
			l.Fatal(err)
		}
		mode = _mode
		// Modes:
		// 0: Return snippet of text at match point, no need to scroll
		// 1: Return whole file
	}
	prevLines := 6
	afterLines := 60
	//Printf("ln %v\nmatch %v\nfile %v\n", ln, match, filePath)

	if fs, err := os.Stat(filePath); err != nil || fs.IsDir() {
		absPath := relDir + filePath
		if fs, err := os.Stat(absPath); err == nil && fs.IsDir() == false {
			filePath = absPath
		} else {
			// Println(match)
			l.Fatalln("ntom: File supplied did not exist or is a directory.\n\t relDir: " + relDir + "\n\t filePath: " + filePath)
		}
	}

	if true || mode == 0 { // fzf will still mis-identify the scrolling position with `+{2}-/2` and these extra lines off, so we might as well have them
		Println(filePath)
		topS.Println(match)
		Println()
		//Println() // The color would leak with Printf
	}
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		l.Fatal(err)
	}
	fileLines := strings.Split(string(fileBytes), "\n")
	if ln != 0 {
		ps := ln - prevLines
		if ps < 0 {
			ps = 0
		}
		if mode == 1 {
			ps = 0
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
		if mode == 1 {
			ae = afterLinesTLen
		}
		Println(strings.Join(fileLines[(ln+1):(ae+1)], "\n"))
	}

	os.Exit(0)
}
