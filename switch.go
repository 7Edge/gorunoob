package main

import "fmt"

func main() {
	var grade string
	var marks = 70

	switch marks {
	case 90 :
		grade = "b"
	case 80:
		grade = "c"
	case 70,60:
		grade = "c-"
	default:
		grade = "d"
	}
	fmt.Printf("core is %v", grade)

	switch {
	case grade == "b":
		fmt.Printf("chengji youxiu")
	case grade == "c":
		fmt.Printf("chengji liang")
	case grade == "c-":
		fmt.Printf("chengji jige")
	case  grade == "d":
		fmt.Printf("chengji bujige")
	default:
		fmt.Printf("chengji cha")
	}

}

