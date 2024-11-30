package main

import (
	"fmt"
)


func Greet(){
	fmt.Printf("hello boy\n")
}


func Greetbyname(name,name2 string , age int){
	fmt.Printf("hello %v %v , ur age is %v\n",name,name2,age)
}

func Sum(n1,n2 int) int{
	return n1+n2
}

func SumMulti(n1,n2 int) (int,int){
	return n1+n2,n1*n2
}

func Byname(n1,n2 int)(sum int,multi int){
	sum = n1+n2
	multi = n1*n2
	return
}


func FuncInFunc(a,b int, action func(a,b int) int) int {
	return action(a,b)
}



func main()  {
	fmt.Println("hello world")
	Greet()

	Greetbyname("bektur","moidinov",17)
	fmt.Println(Sum(3,4))
	fmt.Println(SumMulti(4,4))
	fmt.Println(Byname(6,6))
    var multi func(a,b int) int
	multi = func (a,b int) int {return a+b}
	fmt.Println(multi(10,10))
	fmt.Println(FuncInFunc(3,3,Sum))
	Pointers()

}

