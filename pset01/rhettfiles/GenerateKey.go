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


func main() {

	// declaring your secret and public keys

	var sec SecretKey
	var pub PublicKey

	// generate secret key ZeroPre

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


	fmt.Printf("%08b\n", sec.ZeroPre)
	fmt.Printf("%08b\n", sec.OnePre)
	fmt.Printf("%x\n", pub.ZeroHash)
	fmt.Printf("%x\n", pub.OneHash)


	
	
}

func (self Block) Hash() Block {
	return sha256.Sum256(self[:])
}
