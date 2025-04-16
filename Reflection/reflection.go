package reflection

import (
	"fmt"
	"reflect"
)

// Golang challenge: write a function Reflect(x interface{}, fn func(string)) which takes a struct x and
// calls fn for all strings fields found inside.
func Reflect(x any, fn func(string)) {
	reflection := reflect.ValueOf(x)
	for index := range reflection.NumField() {
		field := reflection.Field(index)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Int:
			fn(fmt.Sprintf("%v", field.Int()))
		}
	}
}
