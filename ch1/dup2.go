package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range fileNames {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
	printDups(counts)
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func printDups(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
