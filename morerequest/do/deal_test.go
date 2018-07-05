package do

import "testing"

func BenchmarkDoWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoWork(i)
	}
}
