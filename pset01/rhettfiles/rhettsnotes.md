Lamport signatures -- Rhett Applestone
----------------------

How do they work?

First generate 256 pairs of 256 bit numbers

Arrange like so into a zero row and a one row


0 [] [] [] []  ... (256 numbers in total here)



1 [] [] [] []  ... (256 numbers in total here)


These two "First rows" will be our Private key

Next hash each of the individual 256 bit numbers to create two new
rows under your two secret key rows (any hashing algorithm works, but lets use sha256 in this example)


0 [] [] [] [] ...
  [] [] [] [] ...

1 [] [] [] [] ...
  [] [] [] [] ...


the top row of both the 0 and 1 row is your secret key, and 
the bottom row is your public key.


Next, publish your public key (the second rows of each row) out 
in the wild for anyone to see.

After you publish your secret key it's time to actually sign the message

Say your message is "My name is joe" and you want to prove that you, with your private key,
were the one who signed and sent it.

To do this, first hash your message, ie type in  echo 'My name is joe' | sha256sum    to your computer's terminal

in our case the hash would be fe3fef781cf932cda2b39a4d0ceef7fac3b59b7871d3c12ac4adac091ac0a2f4

This is in hexadecimal however, but for lamport signatures we need to convert this into a string of binary.

If you didnt know, hexadecimal is base 16, and each hex character can be converted into binary (base 2)


for each hex character we are going to create a 4 binary character block which the binary will go in


[Hex to binary table]


Hex     Binary
0 -->   0000
1 -->   0001
2 -->   0010
3 -->   0011
4 -->   0100
5 -->   0101
6 -->   0110
7 -->   0111
8 -->   1000
9 -->   1001
a -->   1010
b -->   1011
c -->   1100
d -->   1101
e -->   1110
f -->   1111

This is why are hex string is 64 characters long in hex, because it is 256 bits long in binary
 
For our string of hashed data "My name is joe" since the first letter is f we would have 

1111 followed by 1110 because the first and second characters are f and e of the string 




If we converted the whole string we would have
 

1111111000111111111011110111100000011100111110010011001011001101
1010001010110011100110100100110100001100111011101111011111111010
1100001110110101100110110111100001110001110100111100000100101010
1100010010101101101011000000100100011010110000001010001011110100


(Sidenote I made a hex --> binary converter if you need it hexconvert.go)

Remember this is the hash of our original message "My name is joe" in 
sha256 converted to binary

Now we have all the required components that we need for a valid 
lamport signature.













