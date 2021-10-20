package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x type is : %T", i)
	case int:
		fmt.Printf("x type is int")
	case string:
		fmt.Printf("x type is string")
	case float64:
		fmt.Printf("x type is float64")
	case bool:
		fmt.Printf("x type is bool")
	case func(int) float64:
		fmt.Printf("x is func(int)float64")
	default:
		fmt.Printf("Uknown type")
	}
}

	
