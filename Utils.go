package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return file
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

func Assert(check bool, msg string) {
	if !check {
		fmt.Println(msg)
		os.Exit(1)
	}
}
