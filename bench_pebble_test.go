package contrast_benchmark

import (
	"github.com/cockroachdb/pebble"
	"math"
	"path/filepath"
	"testing"
)

var pebbledb *pebble.DB

func init() {
	dir := filepath.Join("benchmark", "pebble")
	opts := &pebble.Options{
		BytesPerSync: math.MaxInt,
	}
	var err error
	pebbledb, err = pebble.Open(dir, opts)
	if err != nil {
		panic(err)
	}
}

func Benchmark_PutValue_Pebble(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		pebbledb.Set(GetKey(i), GetValue(), &pebble.WriteOptions{
			Sync: false,
		})
	}
}

func Benchmark_GetValue_Pebble(b *testing.B) {
	for i := 0; i < 500000; i++ {
		pebbledb.Set(GetKey(i), GetValue(), &pebble.WriteOptions{
			Sync: false,
		})
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pebbledb.Get(GetKey(i))
	}
}
