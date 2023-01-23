package main

// the sum of arrays, 使用切片类型。
// 切片类型不会将集合的长度保存到类型中，所以它的尺寸是不固定的。
func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}
