# BRIEF NOTE TO THOSE READING #
---
After embarking on the task of completing Cryptocurrency Engineering And Design MAS.s62
something became very appartent.

1. I needed to learn golang before I proceeded 

I started the course about 21 days ago with much less programming experence, but a book called

*An Introduction to Programming in Go* helped and is free

the book has a lot of flaws and there is a very similar book by O'Riley called

*Introducing Go* which I would recomend

installing godoc, or golang-golang-x-tools bricked my version of go and I had to reinstall

Since that was ~21 days ago I am going to be relearning about lamport signatures 

the best explanation was by someone called Biz Vlogs on youtube 

https://www.youtube.com/watch?v=v3yApr-8IIo

The original paper looks interesting but I havent looked at it more than a skim

*Constructing Digital Signatures from a One Way Function*

http://lamport.azurewebsites.net/pubs/dig-sig.pdf

also checkout https://lamport.org/ 

---
### Lamport Signatures ###

how to make lamport signatures

first generate 256 pairs of random "numbers" I put numbers in quotes here
because these "numbers" wont be what we normally think of as numbers.

each of these "numbers" is a string of individual 256 bits. this can be expressed in
hexadecimal.



Lamport signature roadblocks. If you have a probem, these may be your solutions.

1. if your bytes are showing up as wierd sizes, you are likely not showing 
the leading zeroes, use printf("%08b", <thing to be printed>) 

2. each preimmage is a 2d array/ slice, read up on them. you can access specific
parts with two brackets key.Zeropre[row][column]

3. you may need two loops inside eachother


---

3-23-23

### What I've been able to acomplish so far ###

Since i started working on the project fully after learning the basics of go
four days ago, I have made some good progress. Fisrtly, I was able to complete
the first of the three required functions, being GenerateKey() Sign() Verify().
I have left the Generate key in a file called GenerateKey, if you want to see 
it. 

*How did you create GenerateKey?*

Firstly, I used the crypto/rand package to generate bytes of data and write it to the
ZeroPre, and OnePre of the Secret key i call sec. nested for loops go over each element
in the nested 2d array. After this is done, I look simply at the blocks (256) each row
and hash them which is stored in the corrosponding public hash. I'm not the best programmer
but I was happy with my implementation. Next steps will be to make a signature function that
takes the desired message, hashes it, and then looks at the individual 256 bits of the binary
representation of the hash, and "reveals" the preimmages.



### 3-26-23 Re-Explaining Lamport signatures to myself ###

I dont know about you, but when I'm learning a specific topic, I like to re-explain it to 
myself to 

1. See if I know what I'm talking about
and
2. Get a better sense of direction for what I am going to need to do next

this process makes everything much clearer, and provides a refresh of the mind.
that is what I will be doing below

Lamport signatures

first, Generate 256 pairs of 32 byte numbers, arrange in two rows

row 0 [][][][][][] ... 256

row 1 [][][][][][] ... 256

each [] represents 32 bytes or 256 bits of random data, we will call it a block.

after we have this random data we will hash it and arrange two more rows like so


row 0   [][][][][][] ... 256
Hash 0  [][][][][][] ... 256

row 1   [][][][][][] ... 256
hash 1  [][][][][][] ... 256

the unhashed rows will be out private key, and the hashed rows will be out public key

now comes time to sign out message. to do this we will first need to hash it.

message := "Hello"
hashedmsg := hash(message)

now we have our hashed message, in this example it's going to be sha256.

now we need to look at the individual bits of the hashed message.

say the bits of the hash are 110110 we would need to revial the first block of our private key in row 1
we will also need to revial the second block of our private key in row 1, but we will revial the third block
of our private key in row 0

all done it would look like 

row 0       []    [] ... 256
Hash 0  [][][][][][] ... 256

row 1   [][]  [][]   ... 256
hash 1  [][][][][][] ... 256

we add all of the revealed blocks from row 0 and 1 together to get a complete 256 block long signature .

to verify our signature we need a few things, firstly, we need to take the signarure's first block and hash it
then compare that hash to each of our hash rows.

first declare a new message of []byte

this will be where we put the data to test if the order of revealed preimages matches the data 
of our hashed message

if it matches the 0 row we change nothing 

if it matches the 1 row we add a 1 to the specific spot  

it it matches with neither, say it matched with neither and stop

we repeat this pattern for all 256 blocks until we have a full new message

then we check if the hash of our original message == new messagex




