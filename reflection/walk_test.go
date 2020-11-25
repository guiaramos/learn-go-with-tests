package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Gui"},
			[]string{"Gui"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Gui", "Seoul"},
			[]string{"Gui", "Seoul"},
		}, {
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Gui", 30},
			[]string{"Gui"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
