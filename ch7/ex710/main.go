// “The sort.Interface type can be adapted to other uses. Write a function IsPalindrome(s sort.Interface) bool that reports whether the sequence s is a palindrome, in other words, reversing the sequence would not change it. Assume that the elements at indices i and j are equal if !s.Less(i, j) && !s.Less(j, i).”

package main

import (
	"fmt"
	"sort"
)

type PalindromeChecker []byte

func (x PalindromeChecker) Len() int           { return len(x) }
func (x PalindromeChecker) Less(i, j int) bool { return x[i] < x[j] }
func (x PalindromeChecker) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(PalindromeChecker([]byte("abcdcba"))))
	fmt.Println(IsPalindrome(PalindromeChecker([]byte("abcdecba"))))
}
