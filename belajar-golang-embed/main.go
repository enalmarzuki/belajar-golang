package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed version.txt
var version string

//go:embed 5.jpg
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println("VERSION : ", version)

	err := os.WriteFile("5_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("FILENAME : ", entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(strings.ToUpper(entry.Name())+" : ", string(file))
		}
	}
}
