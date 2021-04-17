//++ SOURCE: https://siongui.github.io/2017/05/06/go-check-if-two-string-are-anagram/
// Check by sorting and compare

package anagram

import (
	"sort"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

// StringToRuneSlice convert a string to a slice of rune
func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func AreAnagram(s1, s2 string) bool {
	var r1 ByRune = StringToRuneSlice(s1)
	var r2 ByRune = StringToRuneSlice(s2)

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
}
