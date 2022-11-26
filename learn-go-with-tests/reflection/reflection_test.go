package reflection

import (
	"reflect"
	"testing"
)

// [2]
/* func TestWalk(t *testing.T) {
	expected := "Jams"
	var got []string

	x := struct {
		Name string
	}{
		expected,
	}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of funcion calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
} */

// [4]
/* func TestWalk(t *testing.T) {
	cases := []struct {
		// test name
		Name string
		// input test
		Input interface{}
		// out expected test
		ExpectedCalls []string
	}{
		{
			// test name
			"struct with one string field",
			// input test
			struct {
				Name string
				City string
				Age  int
			}{
				"Jams",
				"Mexico",
				31,
			},
			// out expected test
			[]string{"Jams", "Mexico"},
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
} */

// [5]
/* func TestWalk(t *testing.T) {
	cases := []struct {
		// test name
		Name string
		// input test
		Input interface{}
		// out expected test
		ExpectedCalls []string
	}{
		{
			// test name
			"struct with one string field",
			// input test
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{
				"Jams",
				struct {
					Age  int
					City string
				}{
					31,
					"Mexico",
				},
			},
			// out expected test
			[]string{"Jams", "Mexico"},
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
} */

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

/* func TestWalk(t *testing.T) {
	cases := []struct {
		// test name
		Name string
		// input test
		Input interface{}
		// out expected test
		ExpectedCalls []string
	}{
		{
			// test name
			"struct with one string field",
			// input test
			Person{
				"Jams",
				Profile{
					31,
					"Mexico",
				},
			},
			// out expected test
			[]string{"Jams", "Mexico"},
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
} */

func TestWalk(t *testing.T) {
	cases := []struct {
		// test name
		Name string
		// input test
		Input interface{}
		// out expected test
		ExpectedCalls []string
	}{
		{
			// test name
			"struct with one string field",
			// input test
			&Person{
				"Jams",
				Profile{
					31,
					"Mexico",
				},
			},
			// out expected test
			[]string{"Jams", "Mexico"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"array",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"maps",
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

	// Por el cambio del orden de las keys, se valida el valor de la keys
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

	// Channels
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Jams"}
			aChannel <- Profile{34, "Jose"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Jams", "Jose"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Jams"}, Profile{34, "Jose"}
		}

		var got []string
		want := []string{"Jams", "Jose"}

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
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
