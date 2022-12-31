package moneyform_test

import (
	"testing"

	"github.com/benjamin-kraatz/moneyform/moneyform"
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

		actual := moneyform.NewMoneyformString(intVal)

		if strVal != actual {
			t.Fatal("failed. for", intVal, "expected", strVal, ", got", actual)
		}
	}
}

func TestStringToInt(t *testing.T) {
	for index, sample := range samplesStr {
		strVal := sample
		intVal := samplesInt[index]

		actual := moneyform.NewMoneyformInt(strVal)

		if intVal != actual {
			t.Fatal("failed. for", strVal, "expected", intVal, ", got", actual)
		}
	}

	// test panic case
	strVal := "0.abc"
	actual := moneyform.NewMoneyformInt(strVal)

	if actual != -1 {
		t.Fatal("Failed. Expected to 'fail' with -1, got", actual)
	}
}
