package currency

import (
	"testing"
)

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
