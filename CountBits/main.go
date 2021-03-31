package main

import (
	"fmt"
)

func main() {
	var input uint32
	fmt.Scanf("%d", &input)
	fmt.Println(countBits(input))

}
func countBits(num uint32) int32 {
	bits := decimalToBinaryStillReverseOrder(int(num))
	return int32(countNumberOne(bits))
}

func decimalToBinaryStillReverseOrder(n int) []int {
	var binary []int
	x := n
	for x > 0 {
		binary = append(binary, x%2)
		x = x / 2
	}
	// fmt.Println(binary)
	return binary
}

func countNumberOne(bits []int) int {
	counter := 0
	for _, bit := range bits {
		if bit == 1 {
			counter++
		}
	}
	return counter
}
