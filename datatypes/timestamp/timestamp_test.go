package timestamp_test

import (
	"encoding/json"
	"testing"
	"time"

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

func TestFormat(t *testing.T) {
	ts := timeStrDirect{
		FirstSeen: timestamp.New(time.Date(2021, 01, 01, 12, 0, 0, 0,
			time.FixedZone("test", 3600))),
	}

	encoded, err := json.Marshal(&ts)
	if err != nil {
		t.Fatal(err)
	}

	const expected = `{"first_seen":"2021-01-01T11:00:00.000000Z"}`
	if string(encoded) != expected {
		t.Fatalf("Encoded timestamp has wrong format: %s !=%s", string(encoded), expected)
	}
}
