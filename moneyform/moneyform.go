package moneyform

import (
	"fmt"
	"strconv"
	"strings"
)

// NewMoneyformString converts an integer Moneyform into a string representation.
//
// Example: 14112208 returns "141122.08" and any appendix.
// This could be the currency symbol, for example.
func NewMoneyformString(cents int, appendix string) string {
	fullFloated := float32(cents) / 100
	return fmt.Sprintf("%.2f%s", fullFloated, appendix)
}

// Converts a Moneyform string into an integer.
// For example, "99.34 €" results in 9934
// Before any conversion, the string is being cleaned. See CleanupString.
// Note that currency symbols cannot be removed; do this on your own before.
//
// If conversion fails, it returns -1 and an error
func NewMoneyformInt(money string) (int, error) {
	cMoney := CleanupString(money)
	intVal, err := strconv.Atoi(cMoney)

	if err != nil {
		return -1, err
	}

	return intVal, nil
}

// Cleans up a Moneyform string from special characters.
// Note: currency symbols cannot be removed; do this on your own before.
func CleanupString(money string) string {
	var clp = money
	clp = strings.TrimSpace(clp)
	clp = strings.ReplaceAll(clp, ".", "")
	clp = strings.ReplaceAll(clp, ",", "")

	return strings.TrimSpace(clp)
}
