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

func TestBackspaceCompare(t *testing.T) {
	vals := [5][2]string{{"ab#d", "ad#d"}, {"ab##", "c#d#"}, {"a##c", "#a#c"}, {"a#c", "b"}, {"nzp#o#g", "b#nzp#o#g"}}
	expected := [5]bool{true, true, true, false, true}

	for i := 0; i < len(vals); i++ {
		if expected[i] != backspaceCompare(vals[i][0], vals[i][1]) {
			t.Error(vals[i][0], "and", vals[i][1], "match was", !expected[i], "expected", expected[i])
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	expected := []int{1, 2}

	result := removeDuplicates(input)
	if len(expected) != result {
		t.Error("incorrect result length, expected", len(expected), "but got", result)
	}

	for i := 0; i < len(expected); i++ {
		if input[i] != expected[i] {
			t.Error("incorrectly sorted at index", i, "expected", expected[i], "but got", input[i])
		}
	}

	input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected = []int{0, 1, 2, 3, 4}
	result = removeDuplicates(input)
	if len(expected) != result {
		t.Error("incorrect result length, expected", len(expected), "but got", result)
	}

	for i := 0; i < len(expected); i++ {
		if input[i] != expected[i] {
			t.Error("incorrectly sorted at index", i, "expected", expected[i], "but got", input[i])
		}
	}

}
