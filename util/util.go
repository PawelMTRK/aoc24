package util

import "strings"

// Get n-th occurence of substr in string s. For n = 0, returns whole string, for subsequent
// values, it returns n-th index of the substr.
func IndexN(s string, substr string, n int) int {
	i, last := 0, 0
	for range n {
		i = strings.Index(s[last:], substr)
		if i == -1 {
			return -1
		}
		last += i + 1
	}
	// last variable contains the last index of the substring + 1
	// so we decrement it
	i = last - 1
	if i < 0 {
		i = 0
	}
	return i
}

// Get slice of indexes of the given substr in a string s.
func GetIndexes(s string, substr string) []int {
	count := strings.Count(s, substr)
	indexes := make([]int, count)
	if count == 0 {
		return nil
	}
	for i := range count {
		indexes[i] = IndexN(s, substr, i+1)
	}
	return indexes
}

// Only toggles val when appropriate argument is true.
func ToggleBool(val *bool, enable, disable bool) {
	if enable {
		*val = true
	} else if disable {
		*val = false
	}
}

// Compare three values at once
func Compare3(v1, v2, v3 any) bool {
	if v1 == v2 && v2 == v3 {
		return true
	} else {
		return false
	}
}

// Returns true if any of the arguments passed to it is true
func Any(bools ...bool) bool {
	for _, b := range bools {
		if b {
			return true
		}
	}
	return false
}
