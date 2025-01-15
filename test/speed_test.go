package test

import (
	"fmt"
	"testing"

	"github.com/fluffy-melli/korcen-go"
)

func Benchmark0(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		korcen.Check("")
	}
}

func Test0(t *testing.T) {
	fmt.Println(korcen.Check("gpdmdmd"))
}
