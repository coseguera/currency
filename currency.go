package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type Currency int64

func ToCurrency(f float64) Currency {
	half := 0.5
	if f < 0 {
		half -= 1
	}

	return Currency((f * 100) + half)
}

func ParseCurrency(s string) (Currency, error) {
	if s == "" {
		return Currency(0), nil
	}

	s = strings.Replace(s, ",", "", -1)

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return Currency(0), err
	}

	result := ToCurrency(val)

	return result, nil
}

func (c Currency) Float64() float64 {
	return float64(c) / 100
}

func (c Currency) Multiply(f float64) Currency {
	half := 0.5
	if (f < 0) != (c < 0) {
		half -= 1
	}

	return Currency((float64(c) * f) + half)
}

func Divide(c Currency, f float64) Currency {
	half := 0.5
	if (f < 0) != (c < 0) {
		half -= 1
	}

	return Currency((float64(c) / f) + half)
}

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", c.Float64())
}
