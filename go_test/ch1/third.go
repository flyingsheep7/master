package main

import (
	"flag"
	"os"

	//1
	"fmt"
)

var name string
var cmdLine = flag.NewFlagSet("QUESTION", flag.ExitOnError)

func init() {
	//2
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE OF %s:\n", "question[xw-test]")
		flag.PrintDefaults()
	}
	cmdLine.StringVar(&name, "name", "everyone", "xw-object.")

}
func main() {
	//3
	/*
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "USAGE OF %s:\n", "question")
		}
	*/
	flag.Parse()

	fmt.Printf("hello, %s!\n", name)

}

// go run first.go -name=sdf
