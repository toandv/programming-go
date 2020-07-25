package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
)

func main() {
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    for i := 0; i < 1000; i++ {
      go fetch(url, ch)
    }
  }
  for range os.Args[1:] {
    for i := 0; i < 1000; i++ {
      fmt.Println(<-ch)
    }
  }
  elapsed := time.Since(start).Seconds()
  fmt.Printf("Total %0.2fs \n", elapsed)
}

func fetch(url string, ch chan<- string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprintf("Getting %s ended with %v", url, err)
  }
  nBytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close()
  if err != nil {
    ch <- fmt.Sprintf("Reading body of %s ended with %v", url, err)
  }
  elapsedSecs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%0.2fs\t%7d\t%s", elapsedSecs, nBytes, url)
}
