package user

import (
	"github.com/twhiston/factd/lib/plugin"
	"testing"
)

func TestImplements(t *testing.T) {
	var _ plugin.Plugin = (*User)(nil)
}

func BenchmarkFactCollection(b *testing.B) {
	p := User{}
	for n := 0; n < b.N; n++ {
		_, err := p.Facts()
		if err != nil {
			b.Error(err)
		}
	}
}
