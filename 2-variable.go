package main

import "fmt"

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false //関数内外両方で宣言が可能
)

func sub() {
	xi := 1
	xi = 2 //上書き
	// xi := 2 //と再宣言するとエラーになる
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf) //関数内でのみ有効な書き方
	fmt.Printf("%T\n", xf64)          //型を調べてプリントしてくれる
	fmt.Printf("%T\n", xi)            //\nで出力の際に改行してくれる
}

func main() {
	fmt.Println(i, f64, s, t, f)
	sub()
}
