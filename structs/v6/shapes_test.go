package main

import "testing"

// 表格驱动测试 https://github.com/golang/go/wiki/TableDrivenTests
func TestArea(t *testing.T) {
	// 匿名结构体
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, at := range areaTests {
		got := at.shape.Area()
		if got != at.want {
			t.Errorf("got %.2f, but want %.2f", got, at.want)
		}
	}

}
