package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	files, _ := filepath.Glob("*.html")

	for _, filename := range files {
		fileReplace(filename, "anpanman@baikin.co.jp", "jamuojisan@koujixou.com")
	}
}

func fileReplace(filename, src, dst string) {
	bytes, _ := ioutil.ReadFile(filename)
	lines := strings.Replace(string(bytes), src, dst, -1)
	result := []byte(lines)
	ioutil.WriteFile(filename, []byte(result), 0666)
	fmt.Println("The string replaced in this file", filename)
}
