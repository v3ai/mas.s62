package main

import (
	"fmt"
	"crypto/rand"
//	"bytes"
	"log"
)


func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

type Block [1]byte


type PubKey struct{
	ZeroPre [10]Block
}

func main() {

	for {

		rb, err := generateRandomBytes(1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%b\n", rb)
	}
	var key PubKey
//	fmt.Println(key,"\n")

	for i := 0; i < len(key.ZeroPre); i++ {
	
		randbyte, err := generateRandomBytes(2)
		if err != nil {
			log.Fatal(err)
		}
		
		key.ZeroPre[i] = Block(randbyte)

		// fmt.Printf("%b", randbyte)
		
	}

//	fmt.Printf("%b", key.ZeroPre)

}
