package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var arr = make([][]int, 3)
	var arr2 = make([][]int, 3)
	for i := range arr {
		arr[i] = make([]int, 3)
	}
	for i := range arr {
		arr2[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arr[i][j] = rand.Intn(100)
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arr2[i][j] = rand.Intn(100)
		}
	}
	fmt.Println("rand arr   ", arr)
	fmt.Println("rand arr2  ", arr2)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arr[i][j] = arr[i][j] * arr2[i][j]
		}
	}
	fmt.Println()
	fmt.Println("multiplication  ", arr)
}
