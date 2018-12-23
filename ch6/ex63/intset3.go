// (*IntSet).UnionWith computes the union of two sets using |, the word-parallel bitwise OR operator. Implement methods for IntersectWith, DifferenceWith, and SymmetricDifference for the corresponding set operations. (The symmetric difference of two sets contains the elements present in one set or the other but not both.)”

package inset3

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func (s *IntSet) DifferentWith(t *IntSet) {
	temp := s.Copy()
	temp.IntersectWith(t)

	for i := range s.words {
		s.words[i] ^= temp.words[i]
	}
}

func (s *IntSet) DifferentWith1(t *IntSet) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		if i < len(t.words) {
			for j := 0; j < 64; j++ {
				if word&(1<<uint(j))&t.words[i]&(1<<uint(j)) != 0 {
					// Clear the Set bit
					s.words[i] = s.words[i] &^ (1 << uint(j))
				}
			}
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	temp := s.Copy()
	temp.IntersectWith(t)
	s.UnionWith(t)
	s.DifferentWith(temp)
}

// TODO: try to avoid complicate if statement
func (s *IntSet) SymmetricDifference1(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			for j := 0; j < 64; j++ {
				if s.words[i]&(1<<uint(j)) != 0 {
					if word&(1<<uint(j)) == 0 {
						continue
					} else {
						// Clear the Set bit
						s.words[i] = s.words[i] &^ (1 << uint(j))
						continue
					}
				}
				if s.words[i]&(1<<uint(j)) == 0 {
					if word&(1<<uint(j)) == 0 {
						continue
					} else {
						// Set the Set bit
						s.words[i] |= 1 << uint(j)
						continue
					}
				}
			}
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = make([]uint64, len(s.words))
	copy(t.words, s.words)
	return &t
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
