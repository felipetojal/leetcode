package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	
	splitRoot := strings.Split(strs[0], "")
	prefix := splitRoot
	if len(prefix) == 0 {
		return ""
	}

	// Iterating over all strings
	for i := 0; i < len(strs); i++ {
		split := strings.Split(strs[i], "")
		// Iterating over the characters
		min := min(len(split), len(prefix))
		prefix = prefix[:min]
		for j := 0; j < min; j++ {
			comparison := split[j] == prefix[j]
			if !comparison {
				prefix = prefix[:j]
				break
			}
		}

		if len(prefix) == 0 {
			return ""
		}
	}

	return strings.Join(prefix, "")
}

func main() {
	strs := []string{"ab", "a"}
	result := longestCommonPrefix(strs)
	fmt.Printf("Result: %v\n", result)

}
