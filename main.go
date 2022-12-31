package main

import (
	"fmt"

	"github.com/benjamin-kraatz/moneyform/moneyform"
)

func main() {
	sample1Int := moneyform.NewMoneyformInt("2208")
	sample1Str := moneyform.NewMoneyformString(sample1Int)

	fmt.Println("Sample 1 - str:", sample1Str, "-", sample1Int)

	sample2Str := moneyform.NewMoneyformString(8)
	sample2Int := moneyform.NewMoneyformInt("117")
	sample2IntStr := moneyform.NewMoneyformString(sample2Int)

	fmt.Println("Sample 2 - str:", "8 -", sample2Str, "|", sample2Int, "-", sample2IntStr)
}
