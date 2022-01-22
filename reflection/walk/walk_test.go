package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	channel := make(chan Profile)
	go func() {
		channel <- Profile{33, "Berlin"}
		channel <- Profile{34, "Katowice"}
		close(channel)
	}()
	function := func() (Profile, Profile) {
		return Profile{33, "Berlin"}, Profile{34, "Katowice"}
	}
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"multi-dimensional struct",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"struct pointer",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"slices of structs",
			[]Profile{{33, "London"}, {34, "Reykjavik"}},
			[]string{"London", "Reykjavik"},
		},
		{
			"arrays of structs",
			[2]Profile{{33, "London"}, {34, "Reykjavik"}},
			[]string{"London", "Reykjavik"},
		},
		{
			"maps of structs",
			map[string]string{"Foo": "Bar", "Baz": "Boz"},
			[]string{"Bar", "Boz"},
		},
		{
			"channel receiving structs",
			channel,
			[]string{"Berlin", "Katowice"},
		},
		{
			"function returning structs",
			function,
			[]string{"Berlin", "Katowice"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got=%v want=%v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
