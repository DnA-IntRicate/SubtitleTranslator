package main

import (
	"fmt"
	"os"
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

func Assert(check bool, msg string) {
	if !check {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func AssertError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func CreateFile(path string) *os.File {
	f, err := os.Create(path)
	AssertError(err)

	return f
}

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	AssertError(err)

	return file
}
