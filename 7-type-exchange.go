package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x) //floatに型変換
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y) //intに型変換
	fmt.Printf("%T, %v, %d\n", yy, yy, yy)

	var s string = "14"
	// z = int(s) stringからintへの変換はこの書き方ではできない
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T, %v\n", i, i)

	h := "hello world"
	fmt.Println(h[0])
	fmt.Println(string(h[0]))
}
