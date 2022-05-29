package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/wontw/sorting/heap"
)

type (
	Sorter interface {
		Sort(src []int) []int
	}
)

func main() {
	src := getRandomSliceOfInts(20)
	fmt.Println("Source array:\t", src)

	heapRes := heap.Sort(src)

	fmt.Printf("Heap sort res:\t%v\n", heapRes)
}

func getRandomSliceOfInts(len int) []int {
	rand.Seed(time.Now().UnixNano())

	res := make([]int, len)

	for i := 0; i < len; i++ {
		res[i] = rand.Intn(100)
	}

	return res
}
