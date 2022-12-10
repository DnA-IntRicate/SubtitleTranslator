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

func main() {
	t := translator.New()

	fmt.Println("Welcome")
	input := bufio.NewReader(os.Stdin)

	fmt.Print("Input source file: ")
	filePath, _ := input.ReadString('\n')

	//	file, _ := os.Open("Assets/KO 105 tr.srt")
	file, _ := os.Open(filePath)

	fmt.Print("Input destination file: ")
	outPath, _ := input.ReadString('\n')
	outFile := CreateFile(outPath)
	//outFile := CreateFile("Assets/Translated-eng - 105.srt")

	scanner := bufio.NewScanner(file)
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
		fmt.Print(outBuf)
	}

	file.Close()
	outFile.Close()

	fmt.Println("\nTranslation has concluded!")
	fmt.Scanln()
}
