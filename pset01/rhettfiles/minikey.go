package main

import (
	"fmt"
	"crypto/rand"
)

type Block [2]byte

type Key struct {
	ZeroPre [5]Block
}

func main() {

	b := make([]byte,1)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error")
		return
	}
	
	var k Key
	fmt.Println(k)

	for i := 0; i < len(k.ZeroPre); i++ {
		for j := 0; j < len(k.ZeroPre.Block); j++ {
			b := make([]byte,1)
			_, err := rand.Read(b)
			if err != nil {
				fmt.Println("error")
				return
			}
	
			k.ZeroPre[i].Block[j] = b
		}

	
		k.ZeroPre[i] = b
	}
	fmt.Printf("%08b\n", k)
}
