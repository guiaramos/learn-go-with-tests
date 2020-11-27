package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"Nested fields",
			Person{
				"Gui",
				Profile{27, "Seoul"},
			},
			[]string{"Gui", "Seoul"},
		}, {
			"Pointers to things",
			&Person{
				"Gui",
				Profile{27, "Seoul"},
			},
			[]string{"Gui", "Seoul"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Seoul"},
				{34, "London"},
			},
			[]string{"Seoul", "London"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Seoul"},
				{34, "London"},
			},
			[]string{"Seoul", "London"},
		}, {
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "London"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "London"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "London"}
		}

		var got []string
		want := []string{"Berlin", "London"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expectred %+v to contain %q but it didn't", haystack, needle)
	}
}
