//Package gointerview is a collection of interview problems solved in go
package gointerview

import (
	"testing"
)

//the following are unit tests for the functions in interviewProblems
//in the same order as the functions
//
func TestShortestToChar(t *testing.T) {
	expected := [12]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}
	var result = ShortestToChar("loveleetcode", 'e')
	for i := 0; i < len(result); i++ {
		if expected[i] != result[i] {
			t.Error("Incorrect value, expected", expected[i], "but got", result[i], "at index", i)
		}
	}
}
