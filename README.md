# Moneyform

Internal project to handle loss-free money calculations.

Normally you would store money values in `float`s. But based on some calculations, you might have something like this:

```go
value := 1143.42 / 99.18
```

where both values represents money values. Guess the output of `value`? It is exactly `11.528735632183908`.
Obviously this is not the best base to calculate prices, for example.


Here comes **Moneyform**. All money values are stored in cents. A formatter in frontend should the ever set the separator at the second-last position, like this: 

- `0,02 €`      for a Moneyform of 2
- `4,00 €`      for a Moneyform of 400
- `0,53 €`      for a Moneyform of 53
- `2208,14 €`   for a Moneyform of 220814