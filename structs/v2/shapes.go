package main

// 定义结构体：长方形（宽和高）
type Rectangle struct {
	width  float64
	height float64
}

// 计算周长
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.width + rectangle.height) * 2

}

// 计算面积
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
