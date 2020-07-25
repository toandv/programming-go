package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[FileLine]int)
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

func countLines(file *os.File, counts map[FileLine]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[FileLine{input.Text(), file.Name()}]++
	}
}

// FileLine contains a Line and the FileName of the file it belongs to
type FileLine struct {
	Line     string
	FileName string
}

func printDups(counts map[FileLine]int) {
	for fileLine, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, fileLine.FileName, fileLine.Line)
		}
	}
}
