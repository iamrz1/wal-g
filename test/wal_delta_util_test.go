package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wal-g/wal-g/internal"
	"github.com/wal-g/wal-g/internal/walparser"
	"github.com/wal-g/wal-g/testtools"
	"testing"
)

const (
	WalgTestDataFolderPath = "./testdata"
	WalFilename            = "00000001000000000000007C"
	LastWalFilename        = "00000001000000000000007F"
	DeltaFilename          = "000000010000000000000070_delta"
	DeltaFilename2         = "0000000300000000000000A0_delta"
)

var TestLocation = *walparser.NewBlockLocation(1, 2, 3, 4)

func TestGetDeltaFileNameFor(t *testing.T) {
	deltaFilename, err := internal.GetDeltaFilenameFor(WalFilename)
	assert.NoError(t, err)
	assert.Equal(t, DeltaFilename, deltaFilename)
}

func TestGetPositionInDelta(t *testing.T) {
	assert.Equal(t, 12, internal.GetPositionInDelta(WalFilename))
}

func assertContainsTestLocation(t *testing.T, storage testtools.InMemoryStorage) {
	storageDeltaFilePath := "bucket/server/wal_005/000000010000000000000070_delta.mock"
	locationBuffer, _ := storage.Load(storageDeltaFilePath)
	reader := internal.NewBlockLocationReader(&locationBuffer.Data)
	location, err := reader.ReadNextLocation()
	assert.NoError(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, TestLocation, *location)
}
