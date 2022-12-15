package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	translator "go-googletrans"
)

const APP_VERSION string = "1.0"

func Help(exitCode int) {
	fmt.Printf("SubtitleTranslator v%s.\n", APP_VERSION)
	fmt.Print("\n")
	fmt.Println("Valid switches:")
	fmt.Println("-i, --in, --input\t\tSpecify the input file path.")
	fmt.Println("-o, --out, --output\t\tSpecify the file path to ouput translated file.")
	fmt.Println("-s, --src, --source\t\tSpecify the source file's language. (Set to 'auto' by default).")
	fmt.Println("-d, --dst, --destination\tSpecify the language to translate to. (Set to 'English (en)' by default).")
	fmt.Println("-q, --quiet\t\t\tDon't output translation results in terminal.")
	fmt.Print("\n")
	fmt.Println("Valid usages:")
	fmt.Println("Convert from any language implicitly to English: 'SubtitleTranslator -i InputFile.srt -o OutputFile.srt'")
	fmt.Println("Convert explicitly from Turkish implicitly to English: 'SubtitleTranslator -i InputFile.srt -o OutputFile.srt -s tr")
	fmt.Println("Convert explicitly from Turkish implicitly to English: 'SubtitleTranslator -i InputFile.srt -o OutputFile.srt -s tr")
	fmt.Println("Convert explicitly from English explicitly to Urdu: 'SubtitleTranslator -i InputFile.srt -o OutputFile.srt -s en -d ur")

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
			case "-d", "--dst", "--destination":
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
