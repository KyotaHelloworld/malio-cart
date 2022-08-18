package main

import "fmt"

func main() {
	check("./", 256)
	check("unzip", 256)

}

func check(str string, rng int) {
	txt := str
	enc := stringToNumber(txt, 0, rng)
	dec := numToString(enc, 0, rng)

	fmt.Printf("plaintx := %v\n", txt)
	fmt.Printf("encrypt := %v\n", enc)
	fmt.Printf("encrypt := %v\n\n", dec)

}
