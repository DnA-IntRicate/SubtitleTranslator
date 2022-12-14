package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	translator "go-googletrans"
)

func Help(exitCode int) {
	// messages
	os.Exit(exitCode)
}

func main() {
	t := translator.New()

	var SourcePath, DestinationPath string
	var DstLanguage string = "en"
	var SrcLanguage string = "auto"

	var Quiet bool = false

	if (len(os.Args) == 2) && ((os.Args[1] == "help") || (os.Args[1] == "?")) {
		Help(0)
		//	} else {
		//		fmt.Println("Unknown argument!\nRun 'SubtitleTranslator help' for a list of valid arguments.")
		//		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		currentArg := os.Args[i]
		if currentArg[0] == '-' {
			var nextArg string
			Assert(i+1 < len(os.Args), "Switch is missing an argument!") // Array indices are zero-based but array lengths are not
			nextArg = os.Args[i+1]

			switch currentArg {
			case "-i", "--in", "--input":
				SourcePath = nextArg
				break
			case "-o", "--out", "--output":
				DestinationPath = nextArg
				break
			case "-s", "--src", "--source":
				SrcLanguage = nextArg
				break
			case "-d", "--dest", "--destination":
				DstLanguage = nextArg
				break
			case "-q", "--quiet":
				Quiet = true
				break
			default:
				fmt.Printf("Bad switch: '%s'", currentArg)
				Help(1)
			}
		}
	}

	// Abstract this out to a separate function
	inFile := OpenFile(SourcePath)
	outFile := CreateFile(DestinationPath)

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		var outBuf string

		if !isNumber(line) && !strings.Contains(line, "-->") && line != "\n" {
			translated, err := t.Translate(line, SrcLanguage, DstLanguage)
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
