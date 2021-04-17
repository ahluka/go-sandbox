package main

import (
	"fmt"
	"flag"
)

var test = flag.String("t", "", "a test flag")

func main() {
	flag.Parse()
	fmt.Println("test: ", *test)
}

