package binary

import (
	"encoding/base64"
	"encoding/json"
)

type Binary []byte

var Encoding = base64.StdEncoding

func (b Binary) MarshalJSON() ([]byte, error) {
	if b == nil {
		return []byte("null"), nil
	}
	encoded := Encoding.EncodeToString(b)
	return json.Marshal(encoded)
}

func (b *Binary) UnmarshalJSON(data []byte) error {
	var str *string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str == nil {
		*b = nil
		return nil
	}

	res, err := Encoding.DecodeString(*str)
	if err != nil {
		return err
	}

	*b = res
	return nil
}
