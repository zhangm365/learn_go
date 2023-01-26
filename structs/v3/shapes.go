package main

// 长方形的数据类型
type Rectangle struct {
	width  float64
	height float64
}

// 圆
type Circle struct {
	Radius float64
}

// 计算周长
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.width + rectangle.height) * 2

}

// 计算面积
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
