package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "xw-object.")
}
func main() {
	//3
	flag.Parse()
	hello(name)

}

// go run first.go -name=sdf
