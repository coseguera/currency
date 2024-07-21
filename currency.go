package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type Currency struct {
	cents int64
}

func NewCurrency(f float64) Currency {
	half := 0.5
	if f < 0 {
		half -= 1
	}

	return Currency{
		cents: int64((f * 100) + half),
	}
}

func ParseCurrency(s string) (Currency, error) {
	if s == "" {
		return Currency{cents: 0}, nil
	}

	s = strings.Replace(s, ",", "", -1)

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return Currency{cents: 0}, err
	}

	return NewCurrency(val), nil
}

func (c Currency) Float64() float64 {
	return float64(c.cents) / 100
}

func (c Currency) Multiply(f float64) Currency {
	half := 0.5
	if (f < 0) != (c.cents < 0) {
		half -= 1
	}

	return Currency{
		cents: int64((float64(c.cents) * f) + half),
	}
}

func (c Currency) Divide(f float64) Currency {
	half := 0.5
	if (f < 0) != (c.cents < 0) {
		half -= 1
	}

	return Currency{
		cents: int64((float64(c.cents) / f) + half),
	}
}

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", c.Float64())
}
