//Package gointerview is a collection of interview problems solved in go
package gointerview

//ShortestToChar is a function that, given a string and a character,
//returns an int arr that shows the distance
//to the nearest instance of that character for each char in string
//we are guaranteed at least a single instance of the requested character
//https://leetcode.com/problems/shortest-distance-to-a-character/
//4 ms, beats 100% of golang submissions
func ShortestToChar(S string, C byte) []int {
	var fromLeft = make([]int, len(S))
	var fromRight = make([]int, len(S))
	var lastLeftMatch = -1
	var lastRightMatch = len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			lastLeftMatch = i
		}
		fromLeft[i] = lastLeftMatch
		var rightIndex = len(S) - 1 - i
		if S[rightIndex] == C {
			lastRightMatch = rightIndex
		}
		fromRight[rightIndex] = lastRightMatch
	}

	var result = make([]int, len(S))
	for j := 0; j < len(S); j++ {
		if fromLeft[j] == -1 {
			result[j] = fromRight[j] - j
		} else if fromRight[j] == len(S) {
			result[j] = j - fromLeft[j]
		} else {
			if j-fromLeft[j] <= fromRight[j]-j {
				result[j] = j - fromLeft[j]
			} else {
				result[j] = fromRight[j] - j
			}
		}
	}
	return result
}
