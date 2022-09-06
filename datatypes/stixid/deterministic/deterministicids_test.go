package deterministic_test

import (
	"testing"

	"github.com/avast/libstix2/datatypes/stixid/deterministic"

	"github.com/avast/libstix2/objects/sco/file"
)

func TestDetermineFileId(t *testing.T) {
	f := file.New()
	f.Name = "test.bin"
	f.Hashes = &file.FileHashes{
		Sha256: "fe90a7e910cb3a4739bed9180e807e93fa70c90f25a8915476f5e4bfbac681db",
	}
	f.Size = 213

	const expected = "file--f58d7146-b3b5-5c4c-a52d-f48bd58cfa46"
	determinedId, err := deterministic.DetermineId(f)
	if err != nil {
		t.Fatal(err)
	}

	if determinedId != expected {
		t.Fatalf("failed determined id, %s != %s", determinedId, expected)
	}
}
