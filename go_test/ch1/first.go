package main

import (
	"flag"
	//1
	"fmt"
)

var name string

func init() {
	//2
	flag.StringVar(&name, "name", "everyone", "xw-object.")
	//flag.StringVar()
	flag.Usage()

}
func main() {
	//3
	flag.Parse()
	flag.Usage()
	flag.Usage()

	fmt.Printf("hello, %s!\n", name)

}

// go run first.go -name=sdf
