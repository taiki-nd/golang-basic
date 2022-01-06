package main

import "fmt"

const Pi = 3.14

const (
	Username = "test-user"
	Password = "test-pass"
)

func main() {
	fmt.Println(Pi, Username, Password)
	//Pi = 3  //constで宣言すると上書きできない
}
