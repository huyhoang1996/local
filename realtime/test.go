package main

import (
	"fmt"
	mRand "math/rand"
)

func main() {
	// a := make([]int, 5)
	// printSlice("a", a)

	// b := make([]int, 0, 5)
	// printSlice("b", b)

	// c := b[:2]
	// printSlice("c", c)

	// d := c[2:5]
	// printSlice("d", d)

	println("GetRandNumber: ", GetRandNumber(3))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var numbers = []rune("0123456789")

// getRandNum returns a random number
func GetRandNumber(size int) string {
	b := make([]rune, size)
	for i := range b {
		randNum := mRand.Intn(len(numbers))
		println("==> ", randNum)
		b[i] = letters[randNum]
	}
	return string(b)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
