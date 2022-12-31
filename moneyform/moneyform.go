package moneyform

import (
	"fmt"
	"strconv"
	"strings"
)

func NewMoneyformString(cents int) string {
	fullFloated := float32(cents) / 100
	return fmt.Sprintf("%.2f €", fullFloated)
}

func NewMoneyformInt(money string) int {
	cMoney := CleanupString(money)
	intVal, err := strconv.Atoi(cMoney)

	if err != nil {
		return -1
	}

	return intVal
}

func CleanupString(money string) string {
	var tipi = money
	tipi = strings.ReplaceAll(tipi, "€", "")
	tipi = strings.ReplaceAll(tipi, ".", "")
	tipi = strings.ReplaceAll(tipi, ",", "")

	return strings.TrimSpace(tipi)
}
