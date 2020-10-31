package kadane

import (
	"fmt"
	"testing"
)

func TestLargestSubArraySum(t *testing.T) {
	var tests = []struct {
		array []int
		largestSum int
}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-1, 5, 2, -2, 1, 3, -4, 2, -5, 6}, 9},
		{[]int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}, 187},
		{[]int{-2, -3, -4}, 0},
		{[]int{0}, 0},
}

for _, test := range tests {
	testname := fmt.Sprintf("%d,%d", test.array, test.largestSum)
	t.Run(testname, func(t *testing.T) {
			result := LargestSubArraySum(test.array)
			if result != test.largestSum {
					t.Errorf("got %d, expected %d", result, test.largestSum)
			}
	})
}
}