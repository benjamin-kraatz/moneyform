package main

import (
	"fmt"
	"strconv"
)

func main() {}

func NewMoneyformString(cents int16) string {
	fullFloated := float32(cents) / 100
	return fmt.Sprintf("%.2f €", fullFloated)
}

func NewMoneyformInt(money string) int16 {
	intVal, err := strconv.Atoi(money)

	if err != nil {
		panic(err)
	}

	return int16(intVal)
}
