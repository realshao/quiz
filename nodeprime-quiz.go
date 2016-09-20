package main

import (
        "fmt"
        "os"
        "bufio"
        "log"
        "math"
        "path/filepath"
) 

func main() {
    //max word length
    var max_word_len int = 0
    var longest_compound_len int = 0
    var debug int = 0

    //map for checking word existence
    m := make(map[string]bool)

    //assuming the input file "word.list" is in the current directy 
    //TODO: let user input file name from stdin
    file_name,_ := filepath.Abs("./word.list")
    input_file, err := os.Open(file_name)
    if err != nil {
        log.Fatal(err)
    }

    //close file after exit from main
    defer input_file.Close()

    //read file line by line, process each word
    scanner := bufio.NewScanner(input_file)
    for scanner.Scan() {
        line := scanner.Text()
        max_word_len = int(math.Max(float64(max_word_len), float64(len(line))))
        m[line] = true
        if debug == 1 {
            fmt.Println(line)
        }
    }

    if debug == 1 {
        fmt.Printf("max word len %d\n", max_word_len)
    }

    //chehck if each word is compound word
    longest_compound_word := make([]byte, max_word_len+1)
    for word := range m {
        segmentable := make([]bool, max_word_len+1)
        word_len := len(word)
        num_sub_words := 0

        //segmentable[i] == true means word[0:i] could be segmented into sub-words
        segmentable[0] = true
        //check if word can be breaked into sub-word in dictionary
        for i := 1; i <= word_len; i++ {
            for j := i-1; j >= 0; j-- {
                if segmentable[j] == true {
                    _, sub_word_exist := m[word[j:i]]
                    if sub_word_exist == true && !(j ==0 && i == word_len){
                        //word is segmentable up to j-1, and word[j:i] is another sub-word in dictionary
                        segmentable[i] = true
                        num_sub_words ++
                        break
                    }
                }

            }

        }

        if num_sub_words >= 1 && segmentable[word_len] {
            if len(word) > longest_compound_len {
                longest_compound_len = word_len

                //will be faster without coping the longest_compound_word
                longest_compound_word = []byte(word[:])
            }
        }
    }

    fmt.Printf("Longest compound word length is %d (%s)\n", longest_compound_len, longest_compound_word)
}
