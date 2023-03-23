package main

import (
	"fmt"
	"crypto/rand"
	"math/big"
)

func main() {
	b := make([]byte, 1)

	for x := 0; x < len(b); x++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(2))
		if err != nil {
			panic(err)
			fmt.Println("error")
		}
		fmt.Println(nBig)
		
		b[x] = 
	}
	fmt.Println(b)

}
