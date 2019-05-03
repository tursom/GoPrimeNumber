package main

import (
	"fmt"
	"time"

	"GoPrimeNumber/PrimeNumber"
)

type _uarr = []uint64

func forEachIndexed(arr []uint64, action func(u uint64, index uint64, size uint64)) {
	size := uint64(len(arr))
	for i := uint64(0); i < size; i++ {
		action(arr[i], i, size)
	}
}

func PrintDecomposition(num uint64) {
	fmt.Print(num, " = ")
	forEachIndexed(PrimeNumber.Decomposition(num), func(u uint64, index uint64, size uint64) {
		if index != size-1 {
			fmt.Print(u, " * ")
		} else {
			fmt.Println(u)
		}
	})
}

func UsingTime(action func()) int64 {
	t1 := time.Now().UnixNano()
	action()
	t2 := time.Now().UnixNano()
	return (t2 - t1) / 1000
}

func main() {
	//fmt.Print(PrimeNumber.Decomposition(9991))
	fmt.Println(UsingTime(func() {
		PrimeNumber.Calc(2000000000)
		//for i := uint64(1); i <= 1000000; i++ {
		//	PrimeNumber.Decomposition(i)
		//}
	}))
}
