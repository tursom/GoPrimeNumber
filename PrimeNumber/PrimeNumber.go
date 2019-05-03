package PrimeNumber

import (
	"math"

	"awesomeProject/BitArray"
)

var _bitArray = BitArray.BitArray(0, true)
var _maxNumber = uint64(math.Sqrt(float64(^uint64(0))))

func CheckedNumber() uint64 {
	return _bitArray.Size() << 1
}

func Get(num uint64) bool {
	switch {
	case num < 2:
		return false
	case num == 2:
		return true
	case num&1 == 0:
		return false
	case num > _maxNumber:
		Calc(num)
		return GetUntilEx(uint64(math.Sqrt(float64(num))), func(u uint64) bool {
			if num%u == 0 {
				return false
			}
			return true
		})
	default:
		Calc(num)
		return _bitArray.Get(num >> 1)
	}
}

func max(a uint64, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func unMax(num uint64) uint64 {
	if num > _maxNumber {
		return _maxNumber
	} else {
		return num
	}
}

func Calc(num uint64) {
	checkedNumber := CheckedNumber()
	if num > checkedNumber && !_bitArray.Resize(unMax(max(num, checkedNumber+(checkedNumber>>1))>>1)) {
		return
	}
	checkedNumber = CheckedNumber()
	sqrtMaxNumber := uint64(math.Sqrt(float64(checkedNumber))) >> 1
	_bitArray.Down(0)
	checkedNumber >>= 1
	for i := uint64(1); i <= sqrtMaxNumber; i++ {
		if _bitArray.Get(i) {
			doubleI := i<<1 + 1
			for j := i + doubleI; j < checkedNumber; j += doubleI {
				_bitArray.Down(j)
			}
		}
	}
}

func ForEach(action func(uint64)) {
	action(2)
	checkedNumber := CheckedNumber()
	for i := uint64(3); i <= checkedNumber; i += 2 {
		if _bitArray.Get(i >> 1) {
			action(i)
		}
	}
}

func GetUntil(max uint64, action func(uint64)) {
	if max < 2 {
		return
	}
	Calc(max)
	action(2)
	for i := uint64(3); i <= max; i += 2 {
		if _bitArray.Get(i >> 1) {
			action(i)
		}
	}
}

func IndexWith(beg uint64, end uint64, action func(uint64)) {
	if beg > end || end < 2 {
		return
	}
	Calc(end)
	if beg <= 2 {
		action(2)
	}
	for i := beg | 1; i <= end; i += 2 {
		if _bitArray.Get(i >> 1) {
			action(i)
		}
	}
}

func GetUntilEx(max uint64, action func(uint64) bool) bool {
	if max < 2 {
		return false
	}
	if !action(2) {
		return false
	}
	for i := uint64(3); i <= max; i += 2 {
		if _bitArray.Get(i>>1) && !action(i) {
			return false
		}
	}
	return true
}

func DoubleDecomposition(num uint64) (a uint64, b uint64) {
	a, b = 1, num
	GetUntilEx(uint64(math.Sqrt(float64(num))), func(u uint64) bool {
		if num%u == 0 {
			a, b = u, num/u
			return false
		}
		return true
	})
	return
}

type Node struct {
	value uint64
	next  *Node
}

func Decomposition(num uint64) []uint64 {
	a, b := DoubleDecomposition(num)
	size := 0
	list := Node{0, nil}
	endNode := &list
	for a != 1 {
		endNode.next = &Node{a, nil}
		endNode = endNode.next
		size++
		a, b = DoubleDecomposition(b)
	}
	endNode.next = &Node{b, nil}
	endNode = endNode.next
	size++

	ret := make([]uint64, size)
	endNode = list.next
	for i := 0; i < size; i++ {
		ret[i] = endNode.value
		endNode = endNode.next
	}
	return ret
}
