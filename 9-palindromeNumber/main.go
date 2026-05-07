package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	split := strings.Split(str, "")
	fmt.Println(split)

	l := len(split)

	end := l - 1
	start := 0

	for start <= end {
		if split[start] != split[end] {
			return false
		}
		fmt.Printf("split[start: %d] = %v\n", start, split[start])
		fmt.Printf("split[end: %d] = %v\n", end, split[end])
		end--
		start++
	}

	return true
}

func main() {
	// numbers := []int{2319, 41934, 159, 01000, 43321, 52458}
	number := 41934

	b := isPalindrome(number)
	fmt.Println(b)
}
