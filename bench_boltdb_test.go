package contrast_benchmark

import (
	"go.etcd.io/bbolt"
	"os"
	"testing"
)

var boltDB *bbolt.DB

func init() {
	opts := bbolt.DefaultOptions
	opts.NoSync = true
	var err error
	_ = os.MkdirAll("benchmark/boltdb", os.ModePerm)
	boltDB, err = bbolt.Open("benchmark/boltdb/bolt.data", 0644, opts)
	if err != nil {
		panic(err)
	}

	boltDB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucket([]byte("test-bucket"))
		if err != nil {
			panic(err)
		}
		return nil
	})

	initBoltDBData()
}

func initBoltDBData() {
	var k int
	for i := 0; i < 5; i++ {
		boltDB.Update(func(tx *bbolt.Tx) error {
			for j := 0; j < 100000; j++ {
				err := tx.Bucket([]byte("test-bucket")).Put(GetKey(k), GetValue())
				if err != nil {
					panic(err)
				}
				k++
			}
			return nil
		})
	}
}

func Benchmark_PutValue_BoltDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		boltDB.Update(func(tx *bbolt.Tx) error {
			err := tx.Bucket([]byte("test-bucket")).Put(GetKey(i), GetValue())
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func Benchmark_GetValue_BoltDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		boltDB.View(func(tx *bbolt.Tx) error {
			tx.Bucket([]byte("test-bucket")).Get(GetKey(i))
			return nil
		})
	}
}
