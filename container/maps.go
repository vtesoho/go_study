package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxlength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("name"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcabc"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
}
