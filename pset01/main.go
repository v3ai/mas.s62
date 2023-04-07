// Problem set 01: Hash based signatures.

// A lot of this lab is set up and templated for you to get used to
// what may be an unfamiliar language (Go).  Go is syntactically
// similar to C / C++ in many ways, including comments.

// In this pset, you need to build a hash based signature system.  We'll use sha256
// as our hash function, and Lamport's simple signature design.

// Currently this compiles but doesn't do much.  You need to implement parts which
// say "your code here".  It also could be useful to make your own functions or
// methods on existing structs, espectially in the forge.go file.

// If you run `go test` and everything passes, you're all set.

// There's probably some way to get it to pass the tests without making an actual
// functioning signature scheme, but I think that would be harder than just doing
// it the right way :)

package main

import (
//	"bytes"
	"crypto/sha256"
//	"encoding/hex"
	"fmt" 
	"crypto/rand"
	"math"
)

func main() {

	// Define your message
	textString := "Did I pass?? 93710237879"
	fmt.Printf("%s\n", textString)

	// convert message into a block
	m := GetMessageFromString(textString)
	fmt.Printf("%x\n", m[:])

	// generate keys
	sec, pub, err := GenerateKey()
	if err != nil {
		panic(err)
	}

	// print pubkey.
	fmt.Printf("pub:\n%s\n", pub.ToHex())

	// sign message
	sig1 := Sign(m, sec)
	fmt.Printf("sig1:\n%s\n", sig1.ToHex())

	// verify signature
	worked := Verify(m, pub, sig1)

	// done
	fmt.Printf("Verify worked? %v\n", worked)

	// Forge signature (re add later possibly?) -Rhett

/*	msgString, sig, err := Forge()
	if err != nil {
		panic(err)
	}

	fmt.Printf("forged msg: %s sig: %s\n", msgString, sig.ToHex())

	return

*/

}

// Signature systems have 3 functions: GenerateKey(), Sign(), and Verify().
// We'll also define the data types: SecretKey, PublicKey, Message, Signature.

// --- Types

// A block of data is always 32 bytes long; we're using sha256 and this
// is the size of both the output (defined by the hash function) and our inputs
type Block [32]byte

type SecretKey struct {
	ZeroPre [256]Block
	OnePre  [256]Block
}

type PublicKey struct {
	ZeroHash [256]Block
	OneHash  [256]Block
}

// --- Methods on PublicKey type

// ToHex gives a hex string for a PublicKey. no newline at the end
func (self PublicKey) ToHex() string {
	// format is zerohash 0...255, onehash 0...255
	var s string
	for _, zero := range self.ZeroHash {
		s += zero.ToHex()
	}
	for _, one := range self.OneHash {
		s += one.ToHex()
	}
	return s
}

// HexToPubkey takes a string from PublicKey.ToHex() and turns it into a pubkey
// will return an error if there are non hex characters or if the lenght is wrong.
// func HexToPubkey(s string) (PublicKey, error) {
	// var p PublicKey
// 
	// expectedLength := 256 * 2 * 64 // 256 blocks long, 2 rows, 64 hex char per block
// 
	// // first, make sure hex string is of correct length
	// if len(s) != expectedLength {
		// return p, fmt.Errorf(
			// "Pubkey string %d characters, expect %d", expectedLength)
	// }
// 
	// // decode from hex to a byte slice
	// bts, err := hex.DecodeString(s)
	// if err != nil {
		// return p, err
	// }
	// // we already checked the length of the hex string so don't need to re-check
	// buf := bytes.NewBuffer(bts)
// 
	// for i, _ := range p.ZeroHash {
		// p.ZeroHash[i] = BlockFromByteSlice(buf.Next(32))
	// }
	// for i, _ := range p.OneHash {
		// p.OneHash[i] = BlockFromByteSlice(buf.Next(32))
	// }
// 
	// return p, nil
// }

// A message to be signed is just a block.
type Message Block
// this will be a 256 bit number generated from the placement of the 0 and 1 blocks of 
// the signature

// --- Methods on the Block type

// ToHex returns a hex encoded string of the block data, with no newlines.
func (self Block) ToHex() string {
	return fmt.Sprintf("%064x", self[:])
}

// Hash returns the sha256 hash of the block.
func (self Block) Hash() Block {
	return sha256.Sum256(self[:])
}

// IsPreimage returns true if the block is a preimage of the argument.
// For example, if Y = hash(X), then X.IsPreimage(Y) will return true,
// and Y.IsPreimage(X) will return false.
func (self Block) IsPreimage(arg Block) bool {
	return self.Hash() == arg
}

// BlockFromByteSlice returns a block from a variable length byte slice.
// Watch out!  Silently ignores potential errors like the slice being too
// long or too short!
func BlockFromByteSlice(by []byte) Block {
	var bl Block
	copy(bl[:], by)
	return bl
}

// A signature consists of 32 blocks.  It's a selective reveal of the private
// key, according to the bits of the message.
type Signature struct {
	Preimage [256]Block
}

// ToHex returns a hex string of a signature
func (self Signature) ToHex() string {
	var s string
	for _, b := range self.Preimage {
		s += b.ToHex()
	}

	return s
}
// GetMessageFromString returns a Message which is the hash of the given string.
func GetMessageFromString(s string) Message {
	return sha256.Sum256([]byte(s))
}

// --- Functions

// GenerateKey takes no arguments, and returns a keypair and potentially an
// error.  It gets randomness from the OS via crypto/rand
// This can return an error if there is a problem with reading random bytes

/*
For the GenerateKey() Function, what is it exactly that we need to do?

First we need to generate random bytes to put into each block  of the private key

type Block [32]byte

type SecretKey struct {
	ZeroPre [256]Block
	OnePre  [256]Block
}


for this we will use crypto/rand





*/





func GenerateKey() (SecretKey, PublicKey, error) {
	// initialize SecretKey variable 'sec'.  Starts with all 00 bytes.
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

	return sec, pub, nil
}

// Sign takes a message and secret key, and returns a signature.
func Sign(msg Message, sec SecretKey) Signature {
	var sig Signature

	
// go through the individual bytes of the hash of the message
	currentBlock := 0
	for j := 0; j < 32; j++ {
// go through each individual bit of each byte from left to right
		for s := 7; s >= 0; s--{
			if (msg[j] & (1 << s) != 0) {
				sig.Preimage[currentBlock] = sec.OnePre[currentBlock]
			} else {
				sig.Preimage[currentBlock] = sec.ZeroPre[currentBlock]
			}	

			currentBlock++
		}	
	}

	return sig

}

// Verify takes a message, public key and signature, and returns a boolean
// describing the validity of the signature.
func Verify(msg Message, pub PublicKey, sig Signature) bool {
// this will be what should be the hash of your original message
	var gsig Message
	
	for i := 0; i < 256; i++ {

	
		if sig.Preimage[i].Hash() == pub.ZeroHash[i] {
			
		} else if sig.Preimage[i].Hash() == pub.OneHash[i] {
			
			gsig[int64(math.Floor(float64(i/8)))] = gsig[int64(math.Floor(float64(i)/8.0))] | (1 << (7-(i%8)))
		
			
		} else if sig.Preimage[i].Hash() != pub.ZeroHash[i] &&  sig.Preimage[i].Hash() != pub.OneHash[i] {
			return false
		}

	}

	if msg == gsig {
		return true 
	}
	
	return false

}

