package main


import (
	"fmt"
	"crypto/rand"
	"crypto/sha256"
)

type Block [32]byte

type SecretKey struct {
	ZeroPre [2]Block
	OnePre  [2]Block
}

type PublicKey struct {
	ZeroHash [2]Block
	OneHash  [2]Block
}


func main() {

	z := []int{4,2,1,3}
	fmt.Println(z)

	var sk SecretKey
	var pk PublicKey

	// generate secret key ZeroPre

	for i := 0; i < len(sk.ZeroPre); i++ {

	
		for j := 0; j < len(sk.ZeroPre[0]); j++ {
			randbyte := make([]byte, 1)
			_, err := rand.Read(randbyte)
			if err != nil {
				fmt.Println("error")
			}
			sk.ZeroPre[i][j] = randbyte[0] 
		}

		// hash each block 

		hash := sk.ZeroPre[i].Hash()
		pk.ZeroHash[i] = hash
		
	}

	fmt.Printf("%x\n", pk.ZeroHash)

	// Generate secretkey for OnePre

	for i := 0; i < len(sk.OnePre); i++ {
		for j := 0; j < len(sk.OnePre[0]); j++ {
			randbyte := make([]byte, 1)
			_, err := rand.Read(randbyte)
			if err != nil {
				fmt.Println("error")
			}
			sk.OnePre[i][j] = randbyte[0]	
		}
	}


	fmt.Printf("%08b\n", sk.ZeroPre)
	fmt.Printf("%08b", sk.OnePre)
	
	
}

func (self Block) Hash() Block {
	return sha256.Sum256(self[:])
}
