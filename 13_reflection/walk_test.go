package walk

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
			"Struct with One field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},

		{
			"Struct with Two field",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},

		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},

		{
			"Nested fields",
			Person{
				"Yasar",
				Profile{24, "Pune"},
			},
			[]string{"Yasar", "Pune"},
		},

		{
			"Pointer to things",
			&Person{
				"Yasar",
				Profile{24, "Pune"},
			},
			[]string{"Yasar", "Pune"},
		},

		{
			"Slices",
			[]Profile{
				{24, "London"},
				{23, "Pune"},
			},
			[]string{"London", "Pune"},
		},

		{
			"Array",
			[2]Profile{
				{24, "London"},
				{23, "Pune"},
			},
			[]string{"London", "Pune"},
		},

		{
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

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}

		})
	}

}
