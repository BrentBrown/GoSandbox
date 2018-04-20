package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, names)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\n", line)
			fmt.Printf("%s\n\n", names[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, names map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		value := names[input.Text()]
		if !contains(value, f.Name()) {
			names[input.Text()] = append(names[input.Text()], f.Name())
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
