package main

import (
	"flag"
	//1
	"fmt"
)

func main() {
	/*
	var name string //1
	flag.StringVar(&name, "name", "everyone", "xw-object.") //2
	*/
	var name = flag.String("name", "everyone", "xw-object.")
	name := flag.String("name", "everyone", "xw-object.")
	flag.Parse()

	//fmt.Printf("hello, %s!\n", name)
	fmt.Printf("hello, %s!\n", *name)


}

// go run first.go -name=sdf
