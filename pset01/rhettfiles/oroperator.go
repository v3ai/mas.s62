package main

import (
	"fmt"
)

func main() {
	x := 110
	 
	fmt.Printf("%08b\n", 110)
	y := x | 1 << 4
	fmt.Printf("%08b\n", y)
}
