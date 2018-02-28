package basic

import (
	"fmt"
	"reflect"
)

func RunReflectTest() {
	var x float64 = 3.4

	fmt.Println("Value : ", reflect.ValueOf(x), reflect.ValueOf(x).Kind())
}
