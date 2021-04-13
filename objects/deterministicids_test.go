package objects_test

import (
	"fmt"
	"testing"

	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/sco/file"
)

func TestDetermineFileId(t *testing.T) {
	f := file.New()
	f.Name = "test.bin"
	f.Hashes = &file.FileHashes{
		Sha256: "fe90a7e910cb3a4739bed9180e807e93fa70c90f25a8915476f5e4bfbac681db",
	}
	f.Size = 213

	fmt.Println(objects.DetermineId(f))
}
