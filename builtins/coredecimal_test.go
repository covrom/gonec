package core

import (
	"testing"

	"github.com/shopspring/decimal"
)

var x decimal.Decimal

func BenchmarkMul(b *testing.B) {
	y := decimal.New(11234, 4)
	for i := 0; i < b.N; i++ {
		x = x.Mul(y)
	}
}
