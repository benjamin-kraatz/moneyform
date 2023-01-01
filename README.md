# Moneyform

Handle loss-free money calculations.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/benjamin-kraatz/moneyform)

Normally you would store money values in `float`s. But based on some calculations, you might have something like this:

```go
value := 1143.42 / 99.18
```

where both values represent money values. Guess the output of `value`? It is exactly `11.528735632183908`. This could be converted into `int`s of course, but it is either rounded or trimmed.

Obviously this is not the best base to calculate prices.


Here comes **Moneyform**. All money values are stored in cents.

- `0,02 €`      for a Moneyform of 2
- `$ 4,00`      for a Moneyform of 400
- `¥ 0,53`      for a Moneyform of 53
- `2208,04 Zl`  for a Moneyform of 220814


## How to use

### Install

Run `go get github.com/benjamin-kraatz/moneyform` in your project. 

Import `github.com/benjamin-kraatz/moneyform` to use this library.

### Money strings to Moneyform integer

To convert a string like "1411.28 €" into a Moneyform representation, just do

```go
money := "1411.28 €"
cleaned := strings.ReplaceAll(money, "€", "")
moneyformInt, err := moneyform.NewMoneyformInt(money)

if err != nil {
    fmt.Println("Could not convert money string:", err)
    return
}

fmt.Println("Moneyform:", moneyformInt)
// Output: 141128
```

Note: make sure to remove all currency symbols and special characters before. Delimiters like `.` and `,` are removed automatically before conversion, so you do not need to remove these.


### Moneyform integers to string

A moneyform like 8501 can be converted into a string representation easily. You can give it an appendix like the currency symbol.

**Note: `.` (period) is used as the delimiter**

```go
moneyformStr := moneyform.NewMoneyformString(8501, " €")

fmt.Println("Money:", moneyformStr)
// Outputs: 85.01 €
```

If you do not want to have an appendix, pass an empty string as the second parameter.


## Contributing

We're looking forward to having people that expand and improve this library. Just clone the code, do your changes and submit a Pull Request! Make sure to follow best practices like documentation, tests and code exec speed.
