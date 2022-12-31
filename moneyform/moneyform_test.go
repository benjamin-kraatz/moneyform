package moneyform_test

import (
	"testing"

	"github.com/benjamin-kraatz/moneyform/moneyform"
)

func TestIntToString(t *testing.T) {
	samplesInt := []int{
		33,
		1411,
		220801,
		6,
	}

	expectedStrings := []string{
		"0.33 €",
		"14.11 €",
		"2208.01 €",
		"0.06 €",
	}

	for index, sample := range samplesInt {
		intVal := sample
		strVal := expectedStrings[index]

		actual := moneyform.NewMoneyformString(intVal)

		if strVal != actual {
			t.Fatal("failed. for", intVal, "expected", strVal, ", got", actual)
		}
	}
}
