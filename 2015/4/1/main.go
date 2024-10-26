package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)
    
var text = "iwrupvqb"

// Where I learned:
// 1. How to work with hashes 
// 2. That string s are compared with ==, unlike in java. You can compare most things
// with ==, even structs, but not slices, maps and functions
// 3. string appears to also be a slice, since you can apply similar functions to it
func main() {
	counter := 1

    for {
		test := getFirstFiveHash(counter)
		fmt.Println(test)
		if test == "00000" {
			return
		}
		counter++
	}
}

func getFirstFiveHash(counter int) string {
	beforeHash := text + strconv.Itoa(counter)
	hash := md5.Sum([]byte(beforeHash))
    hashString := hex.EncodeToString(hash[:])
	return hashString[:5]
}