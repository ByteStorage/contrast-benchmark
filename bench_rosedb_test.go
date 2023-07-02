package contrast_benchmark

import (
	"github.com/rosedblabs/rosedb/v2"
	"path/filepath"
	"testing"
)

var roseDB *rosedb.DB

func init() {
	opts := rosedb.DefaultOptions
	opts.DirPath = filepath.Join("benchmark", "rosedb")
	var err error
	roseDB, err = rosedb.Open(opts)
	if err != nil {
		panic(err)
	}
}

func initRoseDBData() {
	for i := 0; i < 500000; i++ {
		err := roseDB.Put(GetKey(i), GetValue())
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_PutValue_RoseDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := roseDB.Put(GetKey(i), GetValue())
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_GetValue_RoseDB(b *testing.B) {
	initRoseDBData()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := roseDB.Get(GetKey(i))
		if err != nil && err != rosedb.ErrKeyNotFound {
			panic(err)
		}
	}
}
