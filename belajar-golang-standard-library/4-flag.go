package main

import (
	"flag"
	"fmt"
)

func main() {
	var userName *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("username", *userName)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)

	// run project
	// go run 4-flag.go -username=enal -password="rahasia sekali" -host=123.123.123.3 -port=5504
}
