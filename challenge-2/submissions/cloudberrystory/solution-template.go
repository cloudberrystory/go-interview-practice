package main

import (
	"bufio"
	"fmt"
	"os"
// 	"slices"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
    str := []rune(s)
    // slices.Reverse(str)
    
    for i, j := 0, len(str) - 1; i < j; i, j = i+1, j-1 {
        str[i], str[j] = str[j], str[i]
    }
    
    return string(str)
}
