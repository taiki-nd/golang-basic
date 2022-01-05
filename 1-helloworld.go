package main

import (
	"fmt"
	"os/user"
	"time"
)

/*
func init() {
	fmt.Print("Init!")
}
*/

func bazz() {
	fmt.Println("Buzz")
}

func main() {
	//bazz()
	fmt.Println("Hello world!", "test", time.Now())
	fmt.Println(user.Current())
}
