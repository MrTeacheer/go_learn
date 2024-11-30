package main

import "fmt"



func Is3(){
	var n1 int
	fmt.Scanln(&n1)
	switch{
	case n1<3:
		fmt.Println("hohohohoh")
		fallthrough
	case n1>3:
		fmt.Println("hahahaha")
	case n1==3:
		fmt.Println("huhuhuhuhuh")
	}
}


