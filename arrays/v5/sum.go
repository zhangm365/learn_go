package main

func Sum(nums []int) (sum int) {

	for _, num := range nums {
		sum += num
	}
	return

}

/*
计算切片的尾部元素之和。
使用语法 slice[low:high] 获取部分切片
*/
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // 切片：从索引 1 到最后一个元素
			// 在切片末尾追加一个新元素
			sums = append(sums, Sum(tail))
		}

	}

	return sums

}
