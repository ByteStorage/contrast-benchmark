package contrast_benchmark

import (
	"github.com/ByteStorage/FlyDB/config"
	"github.com/ByteStorage/FlyDB/engine"
	"github.com/ByteStorage/FlyDB/flydb"
	_const "github.com/ByteStorage/FlyDB/lib/const"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

var FlyDB *engine.DB
var err error

func init() {
	opts := config.DefaultOptions
	opts.DirPath = filepath.Join("benchmark", "flydb")

	FlyDB, err = flydb.NewFlyDB(opts)
	if err != nil {
		panic(err)
	}
}

func Benchmark_PutValue_FlyDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		err = FlyDB.Put(GetKey(n), GetValue())
		assert.Nil(b, err)
	}
}

func Benchmark_GetValue_FlyDB(b *testing.B) {
	for i := 0; i < 500000; i++ {
		err = FlyDB.Put(GetKey(i), GetValue())
		assert.Nil(b, err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_, err = FlyDB.Get(GetKey(n))
		if err != nil && err != _const.ErrKeyNotFound {
			panic(err)
		}
	}

}
