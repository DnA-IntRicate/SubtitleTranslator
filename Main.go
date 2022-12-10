package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	translator "github.com/Conight/go-googletrans"
)

func isNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func CreateFile(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	return f
}

func DispatchArg(arg string) {
	argSwitch := arg[1:strings.Index(arg, " ")]
	fmt.Println(argSwitch)
}

func GetCommandLineArgs() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		DispatchArg(args[i])
	}
}

func main() {
	t := translator.New()

	var SourcePath, DestinationPath string
	var Quiet bool = false

	for i := 0; i < len(os.Args); i++ {
		currentArg := os.Args[i]
		if currentArg[0] == '-' {
			var nextArg string
			if i+2 > len(os.Args) { // Array indices are zero-based but array lengths are not
				panic("Switch is missing an argument!")
			} else {
				nextArg = os.Args[i+1]
			}

			switch currentArg {
			case "-s", "--src", "--source":
				SourcePath = nextArg
				fmt.Println("source is " + SourcePath) // Temp
			case "-d", "--dest", "--destination":
				DestinationPath = nextArg
				fmt.Println("dst is " + DestinationPath) // Temp
			case "-q", "--quiet":
				Quiet = true
			default:
				panic("Unrecognized switch used!")
			}
		}
	}

	//fmt.Println("Welcome")

	inFile, _ := os.Open(SourcePath)
	outFile := CreateFile(DestinationPath)

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		var outBuf string

		if !isNumber(line) && !strings.Contains(line, "-->") && line != "\n" {
			translated, err := t.Translate(line, "tr", "en")
			if err != nil {
				panic(err)
			}
			outBuf = translated.Text + "\n"
		} else {
			outBuf = line + "\n"
		}

		outFile.WriteString(outBuf)

		if !Quiet {
			fmt.Print(outBuf)
		}
	}

	inFile.Close()
	outFile.Close()

	fmt.Println("\nTranslation has concluded!")
	fmt.Scanln()
}
