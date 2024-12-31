package currency

import (
	"testing"
)

func TestCurrency_Add(t *testing.T) {
	tests := []struct {
		name string
		c1   Currency
		c2   Currency
		want Currency
	}{
		{"Positive + Positive", NewCurrency(1), NewCurrency(2), NewCurrency(3)},
		{"Positive + Negative", NewCurrency(1), NewCurrency(-2), NewCurrency(-1)},
		{"Negative + Positive", NewCurrency(-1), NewCurrency(2), NewCurrency(1)},
		{"Negative + Negative", NewCurrency(-1), NewCurrency(-2), NewCurrency(-3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c1.Add(tt.c2); got != tt.want {
				t.Errorf("Currency.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrency_Subtract(t *testing.T) {
	tests := []struct {
		name string
		c1   Currency
		c2   Currency
		want Currency
	}{
		{"Positive - Positive", NewCurrency(3), NewCurrency(2), NewCurrency(1)},
		{"Positive - Negative", NewCurrency(3), NewCurrency(-2), NewCurrency(5)},
		{"Negative - Positive", NewCurrency(-3), NewCurrency(2), NewCurrency(-5)},
		{"Negative - Negative", NewCurrency(-3), NewCurrency(-2), NewCurrency(-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c1.Subtract(tt.c2); got != tt.want {
				t.Errorf("Currency.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrency_Multiply(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		c    Currency
		args args
		want Currency
	}{
		{"100x2", NewCurrency(1), args{2}, NewCurrency(2)},
		{"100x-2", NewCurrency(1), args{-2}, NewCurrency(-2)},
		{"-100x2", NewCurrency(-1), args{2}, NewCurrency(-2)},
		{"-100x-2", NewCurrency(-1), args{-2}, NewCurrency(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Multiply(tt.args.f); got != tt.want {
				t.Errorf("Currency.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrency_Divide(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		c    Currency
		args args
		want Currency
	}{
		{"100/2", NewCurrency(1), args{2}, NewCurrency(.50)},
		{"100/-2", NewCurrency(1), args{-2}, NewCurrency(-.50)},
		{"-100/2", NewCurrency(-1), args{2}, NewCurrency(-.50)},
		{"-100/-2", NewCurrency(-1), args{-2}, NewCurrency(.50)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Divide(tt.args.f); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrency_String(t *testing.T) {
	tests := []struct {
		name string
		c    Currency
		want string
	}{
		{"Positive", NewCurrency(10.50), "10.50"},
		{"Negative", NewCurrency(-10.50), "-10.50"},
		{"Zero", NewCurrency(0), "0.00"},
		{"Positive with one decimal", NewCurrency(10.5), "10.50"},
		{"Negative with one decimal", NewCurrency(-10.5), "-10.50"},
		{"Millions", NewCurrency(2000000), "2,000,000.00"},
		{"Millions with one decimal", NewCurrency(2000000.5), "2,000,000.50"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Currency.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCurrency(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s       string
		want    Currency
		wantErr bool
	}{
		{"dollar", "1", NewCurrency(1), false},
		{"dollars", "2", NewCurrency(2), false},
		{"dollarsWithCents", "2.38", NewCurrency(2.38), false},
		{"badText", "nope", Currency{}, true},
		{"badCharacter", "n", Currency{}, true},
		{"negativeDollar", "-1", NewCurrency(-1), false},
		{"negativeDollars", "-2", NewCurrency(-2), false},
		{"negativeDollarsWithCents", "-2.38", NewCurrency(-2.38), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ParseCurrency(tt.s)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseCurrency() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseCurrency() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("ParseCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
