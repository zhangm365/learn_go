package main

import "testing"

// 表格驱动测试 https://github.com/golang/go/wiki/TableDrivenTests
func TestArea(t *testing.T) {
	// 匿名结构体
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, at := range areaTests {
		t.Run(at.name, func(t *testing.T) {
			got := at.shape.Area()
			if got != at.want {
				t.Errorf("%#v got %.2f, but want %.2f", at.shape, got, at.want)
			}
		})

	}

}
