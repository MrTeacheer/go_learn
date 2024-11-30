package main

import "fmt"


func Pointers(){
	var n2 *int
	var n1 int = 4
	n2 = &n1
	pointer := &n1
	*n2+=*n2
	fmt.Println(*pointer,*n2)
	TakeP(&n1)
	fmt.Println(*n2,n1)
}

func TakeP(n3 *int){
	 *n3+=*n3
}