package main

import "math"

// 长方形的数据类型
type Rectangle struct {
	width  float64
	height float64
}

// 计算长方形的面积
func (r Rectangle) Area() float64 {

	return r.height * r.width

}

// 圆
type Circle struct {
	Radius float64
}

// 计算圆的面积
func (c Circle) Area() float64 {

	return math.Pi * c.Radius * c.Radius
}

// 接口定义来声明 Shape 类型
type Shape interface {
	Area() float64
}
