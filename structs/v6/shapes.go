package main

import "math"

// 定义结构体，长方形
type Rectangle struct {
	width  float64
	height float64
}

// 定义方法：计算长方形的面积
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

// 定义三角形的结构体
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 1 / 2.0 * t.Base * t.Height
}

// 接口定义来声明 Shape 类型
type Shape interface {
	Area() float64
}
