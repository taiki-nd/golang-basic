package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("hello " + "world")
	fmt.Println("hello world"[0])
	fmt.Println(string("hello world"[0]))

	var s string = "hello world"
	fmt.Println(strings.Replace(s, "h", "x", 1)) //コピーしたものの文字列を置き換えているだけなので
	fmt.Println(s)                               // hello worldが出力される

	s = strings.Replace(s, "h", "x", 1) //変数の上書きをすることで
	fmt.Println(s)                      //xello worldが出力される

	fmt.Println(strings.Contains(s, "world")) //含まれているかどうか

	fmt.Println("test\n" + "test") //改行

	fmt.Println(`point               point
	point
	point`) //そのまま反映される

	//fmt.Println(""") これはエラー
	fmt.Println("\"")
	fmt.Println(`"`)
}
