package main

import "fmt"
import "unsafe"


const (
	Unknown = 0
	Male = 1
	Female = 2
)

const (
	x = 10
	y
	z
)

var (
	i = 11
	j int
)

const (
	o = iota
	p
	q
	r = "haha"
	s
	t = iota
	u
)

const (
	v = "zjq"
	w
)

const (
	jj = "abc"
	ii = len(jj)
	kk = unsafe.Sizeof(jj)
)

func main(){
	const LENGTH int = 20
	const WIDTH int = 30
	var area int
	area = LENGTH * WIDTH
	const a, b, c = 1, "string", true

	fmt.Printf("area %d", area)
	fmt.Println()
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
	fmt.Println(i, j)
	fmt.Println(o,p,q,r,s,t,u)
	fmt.Println(v,w)
	fmt.Println(jj, ii, kk)
}
