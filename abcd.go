package querybuilder

import (
	"fmt"
	"reflect"
)

func Hello() {
	fmt.Println("Hello, World!")
}

type Employee struct {
	id   int
	name string
}

func parseStruct(emp Employee) {
	for i := 0; i < reflect.ValueOf(emp).NumField(); i++ {
		fmt.Println(reflect.ValueOf(emp).Field(i))
	}
}
