package main

import "fmt"

/*
func foo(param1, param2 int) {

}

func main() {
	foo(10, 20)
	//foo(10, 20, 30) パラメーターは２つのためエラーが出る
}
*/

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, params := range params {
		fmt.Println(params)
	}
}

func main() {
	foo()
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...) // f([1, 2, 3]を展開している)
}
