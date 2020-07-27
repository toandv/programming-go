package main

import (
	"fmt"
	"io"

	// "io/ioutil"
	"net/http"
	"os"
	// "reflect"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Getting %s ended with %v \n", url, err)
			os.Exit(1)
		}
		// body, err := ioutil.ReadAll(resp.Body)
		// fmt.Println(reflect.TypeOf(resp.Body))
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading body ended with %v \n", err)
			os.Exit(1)
		}
		// fmt.Printf("Getting %s: \n %s", url, body)
	}

}
