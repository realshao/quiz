Author: Mingchao/Michael Shao
Contact: shaomc@gmail.com
License: GPLv2

To run the program, put "word.list" in current directory where
the "$GOPATH/bin/quiz" command is issued.

Design of the program:
---------------------
1. The program first pre-process the input corpus, get the maximum word length, 
and store each word in a go map(or hash table in other languages), to 
facilitate the compound word check in step 2.
2. Then it iterates each word in the map, check if it is a compound
word by checking if it can be segmented into multiple words in the input
corpus.

Time Complexity
----------------
Time complexity would be N*Average of (word length)^2, where N is number of words
in input.

Space Complexity
----------------
Space complexity would be N, due to the fact that each input word is stored in 
a hash map.



See below for running result on my MacBook(Intel i5 processor, 8G DDR3 memory, SDD)

Mingchaos-MacBook-Pro:quiz shao$ time $GOPATH/bin/quiz
Longest compound word length is 29 (antidisestablishmentarianisms)

real    0m0.517s
user    0m0.490s
sys 0m0.025s
Mingchaos-MacBook-Pro:quiz shao$ ls -l word.list 
-rw-r--r--  1 shao  staff  2715764 Sep 19 14:34 word.list
Mingchaos-MacBook-Pro:quiz shao$ ls -l nodeprime-quiz.go 
-rw-r--r--  1 shao  staff  2366 Sep 19 16:17 nodeprime-quiz.go
Mingchaos-MacBook-Pro:quiz shao$ pwd
/Users/shao/go-workspace/src/github.com/realshao/quiz

