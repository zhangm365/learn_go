package main

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	lens := len(numbersToSum)
	sums = make([]int, lens)

	for i, nums := range numbersToSum {
		sums[i] = Sum(nums)
	}

	return
}
