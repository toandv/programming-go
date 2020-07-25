package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func main() {
  counts := make(map[string]int)
  for _, fileName := range os.Args[1:] {
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
      fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
      continue
    }
    for _, line := range strings.Split(string(data), "\n") {
      counts[line]++
    }
  }
  printDups(counts)
}

func printDups(counts map[string]int) {
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
