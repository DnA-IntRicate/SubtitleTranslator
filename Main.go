package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	translator "go-googletrans"
)

func Help(exitCode int) {
	fmt.Println("Help menu")
	os.Exit(exitCode)
}

func TranslateFile(srcPath, dstPath, srcLang, dstLang string, quiet bool) error {
	t := translator.New()

	inFile := OpenFile(srcPath)
	outFile := CreateFile(dstPath)

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		var outBuf string

		if !isNumber(line) && !strings.Contains(line, "-->") && line != "\n" {
			translated, err := t.Translate(line, srcLang, dstLang)
			if err != nil {
				return err
			}
			outBuf = translated.Text + "\n"
		} else {
			outBuf = line + "\n"
		}

		outFile.WriteString(outBuf)
		if !quiet {
			fmt.Print(outBuf)
		}
	}

	inFile.Close()
	outFile.Close()

	return nil
}

func main() {
	var SourcePath, DestinationPath string
	var DstLanguage string = "en"
	var SrcLanguage string = "auto"
	var Quiet bool = false

	if len(os.Args) <= 1 {
		Help(1)
	} else if (len(os.Args) == 2) && ((os.Args[1] == "help") || (os.Args[1] == "?")) {
		Help(0)
	}

	validArgFound := false
	for i := 1; i < len(os.Args); i++ {
		currentArg := os.Args[i]
		if currentArg[0] == '-' {
			switch currentArg {
			case "-i", "--in", "--input":
				Assert(i+1 < len(os.Args), "Switch "+"\""+currentArg+"\""+" is missing an argument!")
				i++
				SourcePath = os.Args[i]
				validArgFound = true
			case "-o", "--out", "--output":
				Assert(i+1 < len(os.Args), "Switch "+"\""+currentArg+"\""+" is missing an argument!")
				i++
				DestinationPath = os.Args[i]
				validArgFound = true
			case "-s", "--src", "--source":
				Assert(i+1 < len(os.Args), "Switch "+"\""+currentArg+"\""+" is missing an argument!")
				i++
				SrcLanguage = os.Args[i]
				validArgFound = true
			case "-d", "--dest", "--destination":
				Assert(i+1 < len(os.Args), "Switch "+"\""+currentArg+"\""+" is missing an argument!")
				i++
				DstLanguage = os.Args[i]
				validArgFound = true
			case "-q", "--quiet":
				Quiet = true
				validArgFound = true
			default:
				fmt.Printf("Unknown switch: \"%s\".\n", currentArg)
				Help(1)
			}
		}
	}

	if !validArgFound {
		fmt.Println("Unknown command!")
		Help(1)
	}

	if SourcePath == "" {
		fmt.Println("No source path was specified!")
		os.Exit(1)
	}

	if DestinationPath == "" {
		fmt.Println("No destination path was specified!")
		os.Exit(1)
	}

	fmt.Printf("Translating \"%s\" to \"%s\".\n", SourcePath, DestinationPath)
	fmt.Println("============================================================================")
	err := TranslateFile(SourcePath, DestinationPath, SrcLanguage, DstLanguage, Quiet)
	AssertError(err)

	fmt.Println("\nTranslation has concluded!")
	fmt.Scanln()
}
