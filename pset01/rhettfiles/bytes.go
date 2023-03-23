package main

import (
	"fmt"
	"crypto/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("error", err)
			return
		}
			fmt.Printf("%08b\n", b)

	}


}
