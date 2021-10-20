package main

import (
	"fmt"
	"os"
	"io"
)



//point test
type point struct {
	x int
	y int
}

func main() {
	var code int = 123
	var endDate = "2021-10-19"
	var url = "code=%d&endDate=%s"
	var fmtUrl = fmt.Sprintf(url, code, endDate)
	fmt.Println(fmtUrl)

	const name, age = "Kate", 30
	s := fmt.Sprintf("%s is %d years old!", name, age)
	io.WriteString(os.Stdout, s)
	io.WriteString(os.Stdout, "\n")
	fmt.Println(fmt.Sprintf("%b is real value!", age))
	fmt.Println(fmt.Sprintf("%T is real value!", name))
        
	p := point{1,2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%b\n", 12)
	fmt.Printf("%c\n", 12)
	fmt.Printf("%x\n", 12)
	fmt.Printf("%f\n", 3.14)
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("|%6d|%6d|\n", 12, 20)
	fmt.Printf("|%6.1f|%6.1f|\n", 1.22, 2.20)
	fmt.Printf("|%-6.1f|%-6.1f|\n", 1.22, 2.20)
	fmt.Printf("|%-6s|%-6s|\n", "test1", "price")
	
	var s1 = fmt.Sprintf("a %s", "string")

	fmt.Println(s1)
	fmt.Fprintf(os.Stdout, "an %s\n", "error")
}
