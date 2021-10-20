package main

import (
	"fmt"
)

func main() {
	var name string = "runoob"
	fmt.Println(name)
	
	var b, c  int = 10, 30
	fmt.Println(b, c)

	var x string
	var y byte
	var z bool
	fmt.Println(x, y ,z)


	var a1 *int
	var pir *[]int
	var m map[string] int
	var ch chan int
	var f func(string) int
	var e error

	fmt.Println(a1, pir, m , ch, f, e)


}
