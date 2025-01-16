package korcen

import (
	"fmt"
	"testing"
)

func Benchmark0(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Check("")
	}
}

func Test0(t *testing.T) {
	fmt.Println(Check(""))
}
