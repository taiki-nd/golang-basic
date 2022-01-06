package main

import "fmt"

/*
func add(x int, y int) {
	fmt.Println("add function")
	fmt.Println(x + y)
}

func main() {
	add(10, 20)
}
*/

/*
func add(x, y int) int {
	return x + y
}

func main() {
	r := add(10, 20)
	fmt.Println(r)
}
*/

/*
func add(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)
}
*/

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) { //ここでresultを返すことがわかっているので
	result = price * item
	return result //ここのresultは省略することができる
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)
	/*
		f := func() {
			fmt.Println("inner func")
		}
		f()
	*/
	f1 := func(x int) {
		fmt.Println("inner func", x)
	}
	f1(100)

	func(x int) {
		fmt.Println("inner func", x)
	}(200) //省略した書き方
}
