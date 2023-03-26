// it has become apparent that I will need to look at the individual bits in the 
// hash that I am tryint to sign. To do this we will use bitwise operators.
// this is just a test of usage of how I would use bitwise operators to use for 
// the sign function.


package main

import(
	"fmt"
	"crypto/sha256"
)

func main() {
	message := "mesgd"
	h := sha256.New()
	h.Write([]byte(message))
	hash := h.Sum(nil)

	fmt.Println(hash)


	for j := 0; j < 32; j++ {
		for s := 7; s >= 0; s--{
		fmt.Println(hash[j] & (1 << s) != 0)	
		}	
	}

	fmt.Printf("%08b", hash)

	

}
