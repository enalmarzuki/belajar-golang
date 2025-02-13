package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println("VERSION : ", version)
	fmt.Println("VERSION 2 : ", version2)
}

//go:embed 5.jpg
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("5_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println("A :", string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println("B :", string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println("C :", string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("FILENAME : ", entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(strings.ToUpper(entry.Name())+" : ", string(file))
		}
	}
}
