package main

import "fmt"

func init() {
	fmt.Print("Init!")
}

func bazz() {
	fmt.Println("Buzz")
}

func main() {
	bazz()
	fmt.Println("Hello world!")
}
