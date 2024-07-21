package currency

import (
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Currency struct {
	cents    int64
	language language.Tag
}

func NewCurrency(f float64) Currency {
	return NewCurrencyWithLanguage(f, language.English)
}

func NewCurrencyWithLanguage(f float64, lang language.Tag) Currency {
	half := 0.5
	if f < 0 {
		half -= 1
	}

	return Currency{
		cents:    int64((f * 100) + half),
		language: lang,
	}
}

func ParseCurrency(s string) (Currency, error) {
	return ParseCurrencyWithLanguage(s, language.English)
}

func ParseCurrencyWithLanguage(s string, lang language.Tag) (Currency, error) {
	if s == "" {
		return Currency{cents: 0, language: lang}, nil
	}

	s = strings.Replace(s, ",", "", -1)

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return Currency{cents: 0, language: lang}, err
	}

	return NewCurrency(val), nil
}

func (c Currency) Float64() float64 {
	return float64(c.cents) / 100
}

func (c Currency) Add(c2 Currency) Currency {
	return Currency{
		cents:    c.cents + c2.cents,
		language: c.language,
	}
}

func (c Currency) Subtract(c2 Currency) Currency {
	return Currency{
		cents:    c.cents - c2.cents,
		language: c.language,
	}
}

func (c Currency) Multiply(f float64) Currency {
	half := 0.5
	if (f < 0) != (c.cents < 0) {
		half -= 1
	}

	return Currency{
		cents:    int64((float64(c.cents) * f) + half),
		language: c.language,
	}
}

func (c Currency) Divide(f float64) Currency {
	half := 0.5
	if (f < 0) != (c.cents < 0) {
		half -= 1
	}

	return Currency{
		cents:    int64((float64(c.cents) / f) + half),
		language: c.language,
	}
}

func (c Currency) Equal(c2 Currency) bool {
	return c.cents == c2.cents
}

func (c Currency) GreaterThan(c2 Currency) bool {
	return c.cents > c2.cents
}

func (c Currency) LessThan(c2 Currency) bool {
	return c.cents < c2.cents
}

func (c Currency) GreaterThanOrEqual(c2 Currency) bool {
	return c.cents >= c2.cents
}

func (c Currency) LessThanOrEqual(c2 Currency) bool {
	return c.cents <= c2.cents
}

func (c Currency) IsZero() bool {
	return c.cents == 0
}

func (c Currency) IsPositive() bool {
	return c.cents > 0
}

func (c Currency) IsNegative() bool {
	return c.cents < 0
}

func (c Currency) String() string {
	p := message.NewPrinter(c.language)
	return p.Sprintf("$%.2f", c.Float64())
}
