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
