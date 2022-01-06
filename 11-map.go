package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	m["orange"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	var v, ok = m["apple"]
	fmt.Println(v, ok) // => 100 true つまりappleの値は100で値が存在しますよ

	var vv, okok = m["nothing"]
	fmt.Println(vv, okok) // => 0 false

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)
}
