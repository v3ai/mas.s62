/*
What is it that I need to do to implement the Sign() function.

1. first convert the hash of your message to binary and store it as a variable

2. Get the Secret key 




*/





package main


import (
	"fmt"
	"crypto/rand"
	"crypto/sha256"
)

type Block [32]byte

type SecretKey struct {
	ZeroPre [256]Block
	OnePre  [256]Block
}

type PublicKey struct {
	ZeroHash [256]Block
	OneHash  [256]Block
}

type Message Block

type Signature struct {
	Preimage [256]Block
}



func main() {

// declares the message you eventually want to sign
	textString := "1"

// Prints out that message 	
	fmt.Printf("%s\n", textString)

// Gets the hash of original message
	m := GetMessageFromString(textString)

	fmt.Println(m)
	fmt.Printf("%08b\n", m)
	fmt.Printf("%x\n", m)

	
	
}


func (self Block) Hash() Block {
	return sha256.Sum256(self[:])
}

func GetMessageFromString(s string) Message {
	return sha256.Sum256([]byte(s))
}



// func Sign(msg Message, sec, SecretKey) Signature {
// 
	// var sig Signature
	// 
	// 
// 
// 
	// 
// }






func GenerateKey() (SecretKey, PublicKey, error) {
	var sec SecretKey
	var pub PublicKey

	for i := 0; i < len(sec.ZeroPre); i++ {

	
		for j := 0; j < len(sec.ZeroPre[0]); j++ {
			randbyte := make([]byte, 1)
			_, err := rand.Read(randbyte)
			if err != nil {
				fmt.Println("error")
			}
			sec.ZeroPre[i][j] = randbyte[0] 
		}

		// hash each block 

		hash := sec.ZeroPre[i].Hash()
		pub.ZeroHash[i] = hash
		
	}


	// Generate secretkey for OnePre

	for i := 0; i < len(sec.OnePre); i++ {
		for j := 0; j < len(sec.OnePre[0]); j++ {
			randbyte := make([]byte, 1)
			_, err := rand.Read(randbyte)
			if err != nil {
				fmt.Println("error")
			}
			sec.OnePre[i][j] = randbyte[0]	
		}

		hash := sec.OnePre[i].Hash()
		pub.OneHash[i] = hash
	}

	return sec, pub, nil


	
}
