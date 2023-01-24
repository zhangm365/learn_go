package main

// the sum of arrays
func Sum(nums [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += nums[i]
	}
	return
}
