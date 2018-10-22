package color_test

import (
	"testing"

	. "github.com/gbrlsnchs/color"
)

func BenchmarkText(b *testing.B) {
	red := New(CodeFgRed)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = red.Wrap("Hello, world!")
	}
}
