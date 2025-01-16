package main

import (
	"fmt"
	"reflect"
	_ "reflect"
)

// Union solution
// 0 -> string
// 1 -> int32
// 2 -> int64
// 3 -> float32
// 4 -> float64
type hawktuaunion struct {
	type_   *int
	string  *string
	int32   *int32
	int64   *int64
	float32 *float32
	float64 *float64
}

// Type-Reflection solution
func main() {
	a := 69
	b := "nice"

	k_1 := reflect.ValueOf(a).Kind()
	k_2 := reflect.ValueOf(b).Kind()

	fmt.Println(k_1)
	fmt.Println(k_2)

}
