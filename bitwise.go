package main

import (
	"fmt"
)


type Block [2]byte

func main() {
	fmt.Printf("%08b\n", 0x01&0x01)
	fmt.Printf("%08b\n", 0x10&0x01)

	var m Block
	var n Block
	
	fmt.Printf("%08b\n", m[1]&n[1])
	
}
