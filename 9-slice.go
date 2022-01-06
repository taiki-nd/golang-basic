package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[3])
	fmt.Println(n[2:4])
	fmt.Println(n[:3])
	fmt.Println(n[3:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n[2])

	n = append(n, 1000, 2000, 3000, 4000)
	fmt.Println(n)

	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)
}
