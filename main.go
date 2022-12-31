package main

import (
	"fmt"
	"strconv"
)

func main() {
	sample1Int := NewMoneyformInt("2208")
	sample1Str := NewMoneyformString(sample1Int)

	fmt.Println("Sample 1 - str:", sample1Str, "-", sample1Int)

	sample2Str := NewMoneyformString(8)
	sample2Int := NewMoneyformInt("117")
	sample2IntStr := NewMoneyformString(sample2Int)

	fmt.Println("Sample 2 - str:", "8 -", sample2Str, "|", sample2Int, "-", sample2IntStr)
}

func NewMoneyformString(cents int) string {
	fullFloated := float32(cents) / 100
	return fmt.Sprintf("%.2f €", fullFloated)
}

func NewMoneyformInt(money string) int {
	intVal, err := strconv.Atoi(money)

	if err != nil {
		panic(err)
	}

	return intVal
}
