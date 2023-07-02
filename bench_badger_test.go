package contrast_benchmark

import (
	"github.com/dgraph-io/badger/v3"
	"path/filepath"
	"testing"
)

var badgerdb *badger.DB

func init() {
	opts := badger.DefaultOptions(filepath.Join("benchmark", "badger"))
	opts.SyncWrites = false
	badgerdb, err = badger.Open(opts)
	if err != nil {
		panic(err)
	}
}

func Benchmark_PutValue_Badger(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		badgerdb.Update(func(txn *badger.Txn) error {
			return txn.Set(GetKey(i), GetValue())
		})
	}
}

func Benchmark_GetValue_Badger(b *testing.B) {
	for i := 0; i < 500000; i++ {
		badgerdb.Update(func(txn *badger.Txn) error {
			return txn.Set(GetKey(i), GetValue())
		})
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		badgerdb.View(func(txn *badger.Txn) error {
			txn.Get(GetKey(i))
			return nil
		})
	}
}
