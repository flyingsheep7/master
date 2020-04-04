package main

import (
	"flag"
	"os"

	//1
	"fmt"
)

var name string

func init() {
	//2
	flag.StringVar(&name, "name", "everyone", "xw-object.")
	//flag.StringVar()

}
func main() {
	//3
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE OF %s:\n", "question")
	}
	flag.Parse()

	fmt.Printf("hello, %s!\n", name)

}

// go run first.go -name=sdf
