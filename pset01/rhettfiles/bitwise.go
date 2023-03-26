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

	fmt.Println(hash[0])
	fmt.Printf("%08b\n", hash[0])
	

	fmt.Println(hash[0] & (1 << 3) == 1)

	fmt.Println(1 << 1)
}
