package main

import (
	"fmt"
)


func main() {
	ages := map[string]int{}

	ages["lee"] = 32
	ages["sab"] = 18
	ages["mal"] = 16
	ages["aysh"] = 13

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}

