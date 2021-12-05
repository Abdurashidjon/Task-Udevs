package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	arr := make([][]int, 4)
	arr2 := make([][]int, 4)
	for i := 0; i < 4; i++ {
		arr[i] = make([]int, 4)
	}
	for i := 0; i < 4; i++ {
		arr2[i] = make([]int, 4)
	}
	generate(arr)
	generate(arr2)
	fmt.Println(arr)
	fmt.Println(arr2)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			arr[i][j] = arr[i][j] * arr2[i][j]
		}
	}
	fmt.Println("multiplaction  ", arr)
}

func generate(randMatrix [][]int) {
	for i, arr := range randMatrix {
		for j := range arr {
			randMatrix[i][j] = rand.Intn(100)
		}
	}
	for i, arr2 := range randMatrix {
		for j := range arr2 {
			randMatrix[i][j] = rand.Intn(100)
		}
	}
}
