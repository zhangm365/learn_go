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
	sums = make([]int, lens) // 创建切片的时候指定容量

	for i, nums := range numbersToSum {
		sums[i] = Sum(nums)
	}

	return
}
