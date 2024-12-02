package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

var text = "iwrupvqb"

// Nothing new learned, just adjusted the task to get 6 zeros, and to
// start with the solution of the last task, because if the number doesn't
// start with five zeros, it also does not start with 6
func main() {
	counter := 346386

	for {
		test := getFirstFiveHash(counter)
		if test == "000000" {
			fmt.Println(counter)
			return
		}
		counter++
	}
}

func getFirstFiveHash(counter int) string {
	beforeHash := text + strconv.Itoa(counter)
	hash := md5.Sum([]byte(beforeHash))
	hashString := hex.EncodeToString(hash[:])
	return hashString[:6]
}
