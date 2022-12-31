package moneyform

import (
	"fmt"
	"strconv"
)

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
