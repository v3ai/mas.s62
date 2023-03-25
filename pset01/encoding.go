package main

import (
	"fmt"
//	"encoding/hex"
	"crypto/sha256"
)


type Block [32]byte
type Message Block


func main() {
	str := "rhett"
	hash := GetMessageFromString(str)

	fmt.Printf("%b\n",hash)
}

func GetMessageFromString(s string) Message {
	return sha256.Sum256([]byte(s))
	
}
