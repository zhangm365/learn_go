package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"

	// 测试样例
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{expected},
			[]string{expected},
		},

		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},

		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			/*
				if len(got) != 1 {
					t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
				}
			*/
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("want: %s, but got: %s", test.ExpectedCalls, got)
			}
		})
	}

}
