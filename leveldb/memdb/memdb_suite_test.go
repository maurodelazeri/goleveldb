package memdb

import (
	"testing"

	"github.com/maurodelazeri/goleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
