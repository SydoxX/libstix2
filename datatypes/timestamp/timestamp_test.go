package timestamp_test

import (
	"encoding/json"
	"testing"

	"github.com/avast/libstix2/datatypes/timestamp"
)

type timeStrDirect struct {
	FirstSeen timestamp.Timestamp `json:"first_seen,omitempty"`
}

type timeStrPtr struct {
	FirstSeen *timestamp.Timestamp `json:"first_seen,omitempty"`
}

func TestJsonEncode(t *testing.T) {
	var ts timeStrDirect
	encoded, err := json.Marshal(&ts)
	if err == nil {
		t.Fatal("Encoding of zero timestamp is supposed to return an error.")
	}

	var td timeStrPtr
	encoded, err = json.Marshal(&td)
	if err != nil {
		t.Fatal(err)
	}

	if string(encoded) != "{}" {
		t.Fatal("Unexpected encoded value", string(encoded))
	}
}
