package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	for _, f := range GetFiles(".") {
		fmt.Println(f)
	}
}

func GetFiles(dir string) []string {
	var result []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			result = append(result, GetFiles(fpath)...)
			continue
		}
		result = append(result, fpath)
	}
	return result
}