package kadane

func LargestSubArraySum(array []int) int {
		maxSoFar := 0
		maxEndingHere := 0

		for i := 0; i < len(array); i++ {
			maxEndingHere += array[i]
			if maxSoFar < maxEndingHere {
				maxSoFar = maxEndingHere
			}
			if maxEndingHere < 0 {
				maxEndingHere = 0
			}
		}

		return maxSoFar
}
