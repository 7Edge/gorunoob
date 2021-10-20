package main

import "fmt"

var (
	a int
	s string
)

func main(){
	_, y ,z := numbers()
	fmt.Println(y,z)
}

//numbers do
func numbers()(int, int, string){
	a = 10
	s = "sdflj"
	c := 20
	return a, c, s
}
