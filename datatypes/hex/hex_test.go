package hex_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/avast/libstix2/datatypes/hex"
)

type hexStruct struct {
	Value hex.Hex
}

func TestJson(t *testing.T) {
	var err error
	var s hexStruct

	s.Value, err = hex.FromString("ab122133")
	if err != nil {
		t.Fatal(err)
	}

	encoded, err := json.Marshal(&s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))

	var decoded hexStruct
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(decoded.Value, s.Value) {
		t.Fatalf("Values do not match after encoded/decode: %v != %v", s.Value, decoded.Value)
	}
}

func TestJsonNull(t *testing.T) {
	var err error
	var s hexStruct

	encoded, err := json.Marshal(&s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encoded))

	var decoded hexStruct
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}

	if decoded.Value != nil || s.Value != nil {
		t.Fatalf("Values do not match after encoded/decode: %v != %v", s.Value, decoded.Value)
	}
}
