package main

import (
	"fmt"
	"os"
	"strings"
)

func withProgramName() {
	fmt.Println(strings.Join(os.Args, " "))
}

func withIndex() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}

func main() {
	withProgramName()
	withIndex()
}
