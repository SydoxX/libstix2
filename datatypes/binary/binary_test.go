package binary_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/avast/libstix2/datatypes/binary"
)

type binStruct struct {
	Value binary.Binary
}

func TestJson(t *testing.T) {
	var err error
	var s binStruct

	s.Value = binary.Binary{4, 8, 15, 16, 23, 42}

	encoded, err := json.Marshal(&s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))

	var decoded binStruct
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(decoded.Value, s.Value) {
		t.Fatalf("Values do not match after encoded/decode: %v != %v", s.Value, decoded.Value)
	}
}

func TestJsonNull(t *testing.T) {
	var err error
	var s binStruct

	encoded, err := json.Marshal(&s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))

	var decoded binStruct
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}

	if decoded.Value != nil || s.Value != nil {
		t.Fatalf("Values do not match after encoded/decode: %v != %v", s.Value, decoded.Value)
	}
}
