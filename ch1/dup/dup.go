package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, filenames := range counts {
		if len(filenames) > 1 {
			fmt.Printf("%s:\t", strings.Join(dedup(filenames), " "))
			fmt.Printf("%s\n", line)
		}
	}
}

func countLines(f *os.File, counts map[string][]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], filename)
	}
}

func dedup(input []string) []string {
	output := []string{}
	for _, v := range input {
		if !contains(output, v) {
			output = append(output, v)
		}
	}
	return output
}

// from http://stackoverflow.com/a/10485970/3945932
func contains(h []string, n string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}
