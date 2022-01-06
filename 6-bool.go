package main

import "fmt"

func main() {
	//var t, f bool = true, false
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, t, t)
	fmt.Printf("%T %v %t\n", f, f, f)
	fmt.Printf("%T %v %t\n", t, t, 0)
	fmt.Printf("%T %v %t\n", f, f, 1)
	fmt.Printf("%T %v %t\n", t, 0, t)
	fmt.Printf("%T %v %t\n", f, 1, f)
	/*
		%T type
		%v value
		%t bool型かどうか
	*/

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}
