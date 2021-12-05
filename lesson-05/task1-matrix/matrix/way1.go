package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [3][3]int
	var arr2 [3][3]int
	// rand arr
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arr[i][j] = rand.Intn(100)
		}
	}

	//rand arr2
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arr2[i][j] = rand.Intn(100)
		}
	}

	fmt.Println("matrix1  ", arr)
	fmt.Println("matrix2  ", arr2)

	// nxn matrix multiplication
	var multip [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			multip[i][j] = arr[i][j] * arr2[i][j]
		}
	}

	fmt.Println()
	fmt.Println("Multiplication ", multip)

}
