package moneyform_test

import (
	"strings"
	"testing"

	"github.com/benjamin-kraatz/moneyform"
)

var (
	samplesInt []int = []int{
		33,
		1411,
		220801,
		6,
	}

	samplesStr []string = []string{
		"0.33 €",
		"14.11 €",
		"2208.01 €",
		"0.06 €",
	}
)

func TestIntToString(t *testing.T) {
	for index, sample := range samplesInt {
		intVal := sample
		strVal := samplesStr[index]

		actual := moneyform.NewMoneyformString(intVal, " €")

		if strVal != actual {
			t.Fatal("failed. for", intVal, "expected", strVal, ", got", actual)
		}
	}
}

func TestStringToInt(t *testing.T) {
	for index, sample := range samplesStr {
		strVal := sample
		intVal := samplesInt[index]
		strVal = strings.ReplaceAll(strVal, "€", "")

		actual, _ := moneyform.NewMoneyformInt(strVal)

		if intVal != actual {
			t.Fatal("failed. for", strVal, "expected", intVal, ", got", actual)
		}
	}

	// test error case
	strVal := "ab.c"
	actual, err := moneyform.NewMoneyformInt(strVal)

	if err == nil {
		t.Fatal("Failed. Expected to fail with error, got", actual)
	}
}
