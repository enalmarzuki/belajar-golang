// https://github.com/google/re2/wiki/Syntax
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])l`)

	fmt.Println(regex.MatchString("Enal"))
	fmt.Println(regex.MatchString("eNal"))
	fmt.Println(regex.MatchString("eal"))

	fmt.Println(regex.FindAllString("enal eal eNal ENAL, emal", 10))
}
