package decimal

import (
	"testing"
)

var x Decimal

func BenchmarkAdd(b *testing.B) {
	y := New(11234, 4)
	for i := 0; i < b.N; i++ {
		x = x.Add(y)
	}
}
