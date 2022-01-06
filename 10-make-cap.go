package main

import "fmt"

func main() {
	n := make([]int, 3, 5) //make({int, len, cap})
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// => len=3 cap=5 value=[0 0 0]と出力される。sliceにあと2つ余力がある
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// => len=5 cap=5 value=[0 0 0 0 0]
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a) //capを指定しない場合同じになる

	b := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b) //capを指定しない場合同じになる
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c) //capを指定しない場合同じになる

	c = make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	var d []int
	d = make([]int, 0, 5)
	//fmt.Println(d)
	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}
	fmt.Println(d)

}
