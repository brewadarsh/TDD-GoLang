package reflection

import (
	"reflect"
	"testing"
)

func TestReflection(t *testing.T) {
	// The subject to be walked through.
	cases := []struct {
		Name     string
		Input    any
		Expected []string
	}{{
		"struct with one string field",
		struct {
			name string
		}{"name_field"},
		[]string{"name_field"},
	}, {
		"struct with two string field",
		struct {
			name     string
			username string
		}{"name_field", "username_field"},
		[]string{"name_field", "username_field"},
	}, {
		"struct with one string and one int field",
		struct {
			name string
			age  int
		}{"name_field", 25},
		[]string{"name_field", "25"},
	}}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Reflect(test.Input, func(s string) {
				got = append(got, s)
			})
			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("Expected %q but got %q", test.Expected, got)
			}
		})
	}
}
