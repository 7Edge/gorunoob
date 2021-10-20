package main

import (
	"fmt"
	//"unsafe"
	//"io"
	//"os"
)

var str1 = "abc"
var p1str1 = &str1
var p2str1 = &str1
var str2 = str1
var str3 = "xyz"
var p1 = &str1
var p2 = &str2

func main()(){
	if ( p1 == p2 ) {
		fmt.Println("str1 and str2 point equal")
	}
	
	//change 
	*p1str1 = "xyz"
	fmt.Printf("str1 is %s, p1str1 point value is %s , p2str1 point value is %s", str1, *p1str1, *p2str1)
}
