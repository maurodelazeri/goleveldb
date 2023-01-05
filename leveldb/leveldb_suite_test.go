package leveldb

import (
	"testing"

	"github.com/maurodelazeri/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
