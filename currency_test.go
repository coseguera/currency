package currency

import "testing"

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
		{"100x2", Currency(100), args{2}, Currency(200)},
		{"100x-2", Currency(100), args{-2}, Currency(-200)},
		{"-100x2", Currency(-100), args{2}, Currency(-200)},
		{"-100x-2", Currency(-100), args{-2}, Currency(200)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Multiply(tt.args.f); got != tt.want {
				t.Errorf("Currency.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
